package saver

import (
	"context"
	"sync"

	"github.com/ozoncp/ocp-user-api/internal/alarm"
	"github.com/ozoncp/ocp-user-api/internal/flusher"
	"github.com/ozoncp/ocp-user-api/internal/models"
)

type Saver interface {
	Init(ctx context.Context)
	Save(ctx context.Context, user models.User)
	Close()
}

func NewSaver(
	capacity int,
	alarm alarm.Alarm,
	flusher flusher.Flusher,
) Saver {
	return &saver{
		buffer: usersBuffer{
			users: make([]models.User, 0, capacity),
		},
		capacity: capacity,
		done:     make(chan struct{}),
		alarm:    alarm,
		flusher:  flusher,
	}
}

type usersBuffer struct {
	sync.Mutex
	users []models.User
}

// Реализация интерфейса Saver на основе slice. Новые элементы добавляются в конец буфера.
// При заполнении буфера, удаляется первый элемент.
type saver struct {
	buffer   usersBuffer
	capacity int
	done     chan struct{}
	close    chan struct{}
	alarm    alarm.Alarm
	flusher  flusher.Flusher
}

func (s *saver) Init(ctx context.Context) {
	go s.run(ctx)
}

func (s *saver) Save(ctx context.Context, user models.User) {
	s.buffer.Lock()
	defer s.buffer.Unlock()

	if len(s.buffer.users) == s.capacity {
		s.buffer.users = append(s.buffer.users[1:], user)
	} else {
		s.buffer.users = append(s.buffer.users, user)
	}
}

func (s *saver) flush(ctx context.Context) {
	s.buffer.Lock()
	defer s.buffer.Unlock()

	processed := s.flusher.Flush(ctx, s.buffer.users)

	if processed != len(s.buffer.users) {
		s.buffer.users = s.buffer.users[processed:]
	} else {
		s.buffer.users = make([]models.User, 0, s.capacity)
	}
}

func (s *saver) dispose(ctx context.Context) {
	s.flush(ctx)
	s.done <- struct{}{}
}

func (s *saver) run(ctx context.Context) {
	flushSignal := s.alarm.Alarm()

	for {
		select {
		case <-flushSignal:
			s.flush(ctx)

		case <-s.close:
			s.dispose(ctx)
			return

		case <-ctx.Done():
			s.dispose(ctx)
			return
		}
	}
}

func (s *saver) Close() {
	s.close <- struct{}{}
	<-s.done
}
