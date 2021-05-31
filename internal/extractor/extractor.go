package extractor

import (
	"context"
	"github.com/ozoncp/ocp-user-api/internal/models"
	"github.com/ozoncp/ocp-user-api/internal/repo"
)

type NotFoundError string

func (err NotFoundError) Error() string {
	return string(err)
}

const UserNotFoundError NotFoundError = "user not found"

// Результат пакетного извлечуения пользователей из хранилища.
// Поле users содержит значения извлеченных пользователей.
// Поле errors содержит ошибки, которые возникли для конкретных пользователей при извлечении.
type ExtractResult struct {
	Users  []models.User
	Errors map[uint64]error
}

// Интерфейс предназначен для пакетного извлечения пользователей по их идентификаторам из хранилища.
// Не гарантируется порядок идентификаторв, указанных в запросе, и в результате извлечения.
type Extractor interface {
	Extract(ctx context.Context, userIds []uint64) ExtractResult
}

// Значение параметра chunkSize должно быть больше 0 и меньше 100.
// Значение параметра userRepo не может быть nil.
func NewExtractor(
	chunkSize int,
	userRepo repo.Repo,
) Extractor {
	// if chunkSize <= 0 || chunkSize > 100 {
	// 	return nil, errors.New("chunkSize must be in the interval (0, 100]")
	// }

	// if userRepo == nil {
	// 	return nil, errors.New("userRepo can't be nil")
	// }

	return &extractor{
		chunkSize: chunkSize,
		userRepo:  userRepo,
	}
}

type extractor struct {
	chunkSize int
	userRepo  repo.Repo
}

func (e *extractor) Extract(ctx context.Context, usersIds []uint64) ExtractResult {
	users := make([]models.User, 0, len(usersIds))
	var errors map[uint64]error

	if len(usersIds) > 0 {
		for start, end := 0, e.chunkSize; start < len(usersIds); start, end = start+e.chunkSize, end+e.chunkSize {
			if end > len(usersIds) {
				end = len(usersIds)
			}

			chunk := usersIds[start:end]

			if items, err := e.userRepo.GetUsers(ctx, chunk); err != nil {
				if errors == nil {
					errors = make(map[uint64]error)
				}

				for _, userId := range chunk {
					errors[userId] = err
				}
			} else {
				if len(chunk) != len(items) {
					if errors == nil {
						errors = make(map[uint64]error)
					}

					set := map[uint64]struct{}{}

					for _, item := range items {
						set[item.Id] = struct{}{}
					}

					for _, userId := range chunk {
						if _, exists := set[userId]; !exists {
							errors[userId] = UserNotFoundError
						}
					}
				}

				users = append(users, items...)
			}
		}
	}

	return ExtractResult{
		Users:  users,
		Errors: errors,
	}
}
