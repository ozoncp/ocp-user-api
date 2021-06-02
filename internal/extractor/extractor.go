package extractor

import (
	"context"
	"errors"
	"github.com/ozoncp/ocp-user-api/internal/models"
	"github.com/ozoncp/ocp-user-api/internal/repo"
)

var UserNotFoundError = errors.New("user not found")

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

func NewExtractor(
	chunkSize int,
	userRepo repo.Repo,
) Extractor {
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
	errors := make(map[uint64]error)

	chunks := splitArrayToChunks(usersIds, e.chunkSize)

	for _, chunk := range chunks {

		if items, err := e.userRepo.GetUsers(ctx, chunk); err != nil {
			for _, userId := range chunk {
				errors[userId] = err
			}
		} else {
			notFoundUsers := getNotFoundUsers(items, chunk)

			for _, userId := range notFoundUsers {
				errors[userId] = UserNotFoundError
			}

			users = append(users, items...)
		}
	}

	return ExtractResult{
		Users:  users,
		Errors: errors,
	}
}

func getNotFoundUsers(foundUsers []models.User, allUsers []uint64) []uint64 {
	count := len(allUsers) - len(foundUsers)
	result := make([]uint64, 0, count)

	if count > 0 {
		set := map[uint64]struct{}{}

		for _, user := range foundUsers {
			set[user.Id] = struct{}{}
		}

		for _, userId := range allUsers {
			if _, exists := set[userId]; !exists {
				result = append(result, userId)
			}
		}
	}

	return result
}

func splitArrayToChunks(items []uint64, chunkSize int) [][]uint64 {
	result := make([][]uint64, 0, len(items))

	for i, j := 0, chunkSize; i < len(items); i, j = i+chunkSize, j+chunkSize {
		if j > len(items) {
			j = len(items)
		}

		result = append(result, items[i:j])
	}

	return result
}
