package flusher

import (
	"context"
	"github.com/ozoncp/ocp-user-api/internal/models"
	"github.com/ozoncp/ocp-user-api/internal/repo"
	"github.com/ozoncp/ocp-user-api/internal/utils"
)

type Flusher interface {
	Flush(ctx context.Context, users []models.User) int
}

func NewFlusher(
	chunkSize int,
	userRepo repo.Repo,
) Flusher {
	return &flusher{
		chunkSize: chunkSize,
		userRepo:  userRepo,
	}
}

type flusher struct {
	chunkSize int
	userRepo  repo.Repo
}

// Сброс коллекции пользователей в БД.
// Возвращает количество сохраненных сущностей.
func (f *flusher) Flush(ctx context.Context, users []models.User) int {

	chunks, err := utils.SplitToChunks(users, f.chunkSize)

	if err != nil {
		return 0
	}

	for chunkIndex, chunk := range chunks {
		if _, err := f.userRepo.CreateUsers(ctx, chunk); err != nil {
			return chunkIndex * f.chunkSize
		}
	}

	return len(users)
}
