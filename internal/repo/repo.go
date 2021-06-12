package repo

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/ozoncp/ocp-user-api/internal/models"
)

const (
	tableName = "tasks"
)

var (
	errNotImplemented = errors.New("not implemented")
)

type Repo interface {
	CreateUser(ctx context.Context, user *models.User) (uint64, error)
	CreateUsers(ctx context.Context, users []models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
	RemoveUser(ctx context.Context, userId uint64) error
	GetUser(ctx context.Context, userId uint64) (*models.User, error)
	GetUsers(ctx context.Context, userIds []uint64) ([]models.User, error)
	SearchUsers(ctx context.Context, params models.UserSearchParams) (*models.UserSearchResult, error)
}

func NewRepo(
	db *sqlx.DB,
) Repo {
	return &repo{db: db}
}

type repo struct {
	db *sqlx.DB
}

func (r *repo) CreateUser(ctx context.Context, user *models.User) (uint64, error) {
	query := squirrel.Insert(tableName).
		Columns("calendar", "resume", "name", "surname", "patronymic", "email").
		Values(user.CalendarId, user.ResumeId, user.Name, user.Surname, user.Patronymic, user.Email).
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	if query.QueryRow().Scan(&user.Id) != nil {
		return 0, nil
	}

	return user.Id, nil
}

func (r *repo) GetUser(ctx context.Context, userId uint64) (*models.User, error) {
	query := squirrel.Select("id", "calendar", "resume", "name", "surname", "patronymic", "email").
		From(tableName).
		Where(squirrel.Eq{"id": userId}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	var user models.User

	if err := query.QueryRowContext(ctx).Scan(
		&user.Id,
		&user.CalendarId,
		&user.ResumeId,
		&user.Name,
		&user.Surname,
		&user.Patronymic,
		&user.Email,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repo) RemoveUser(ctx context.Context, userId uint64) error {
	query := squirrel.Delete(tableName).
		Where(squirrel.Eq{"id": userId}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	_, err := query.ExecContext(ctx)
	return err
}

func (r *repo) SearchUsers(ctx context.Context, params models.UserSearchParams) (*models.UserSearchResult, error) {
	query := squirrel.Select("id", "calendar", "resume", "name", "surname", "patronymic", "email").
		From(tableName).
		Where(squirrel.Gt{"id": params.Offset}).
		RunWith(r.db).
		Limit(params.Count).
		PlaceholderFormat(squirrel.Dollar)

	rows, err := query.QueryContext(ctx)

	if err != nil {
		return nil, err
	}

	var users []models.User

	for rows.Next() {
		var user models.User

		if err := rows.Scan(
			&user.Id,
			&user.CalendarId,
			&user.ResumeId,
			&user.Name,
			&user.Surname,
			&user.Patronymic,
			&user.Email,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	nextOffset := params.Offset

	if len(users) > 0 {
		nextOffset = users[len(users)-1].Id
	}

	return &models.UserSearchResult{
		Items:      users,
		NextOffset: nextOffset,
	}, errNotImplemented
}

func (r *repo) CreateUsers(ctx context.Context, users []models.User) error {
	return errNotImplemented
}

func (r *repo) UpdateUser(ctx context.Context, user *models.User) error {
	return errNotImplemented
}

func (r *repo) GetUsers(ctx context.Context, userIds []uint64) ([]models.User, error) {
	return nil, errNotImplemented
}
