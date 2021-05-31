package repo

import (
	"context"
	"github.com/ozoncp/ocp-user-api/internal/models"
)

type Repo interface {
	CreateUser(ctx context.Context) (uint64, error)
	UpdateUser(ctx context.Context, user *models.User) error
	RemoveUser(ctx context.Context, userId uint64) error
	GetUser(ctx context.Context, userId uint64) (*models.User, error)
	GetUsers(ctx context.Context, userIds []uint64) ([]models.User, error)
	SearchUsers(ctx context.Context, filter models.UserSearchParams) (models.UserSearchResult, error)
}
