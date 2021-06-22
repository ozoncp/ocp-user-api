package repo

import (
	"context"
	"database/sql"
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
	CreateUsers(ctx context.Context, users []models.User) ([]uint64, error)
	UpdateUser(ctx context.Context, user *models.User) (bool, error)
	RemoveUser(ctx context.Context, userId uint64) (bool, error)
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
	queryBuilder := squirrel.Select("id", "calendar", "resume", "name", "surname", "patronymic", "email").
		From(tableName).
		Where(squirrel.Eq{"id": userId}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var users []*models.User
	if err = r.db.SelectContext(ctx, &users, query, args...); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	var user *models.User
	if len(users) > 0 {
		user = users[0]
	}

	return user, nil
}

func (r *repo) RemoveUser(ctx context.Context, userId uint64) (bool, error) {
	query := squirrel.Delete(tableName).
		Where(squirrel.Eq{"id": userId}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	res, err := query.ExecContext(ctx)
	if err != nil {
		return false, err
	}

	rows, err := res.RowsAffected()
	return rows != 0, err
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

	defer rows.Close()

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

func (r *repo) CreateUsers(ctx context.Context, users []models.User) ([]uint64, error) {
	query := squirrel.Insert(tableName).
		Columns("calendar", "resume", "name", "surname", "patronymic", "email").
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	for _, user := range users {
		query = query.Values(user.CalendarId, user.ResumeId, user.Name, user.Surname, user.Patronymic, user.Email)
	}

	rows, err := query.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	ids := make([]uint64, 0, len(users))
	var id uint64

	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

func (r *repo) UpdateUser(ctx context.Context, user *models.User) (bool, error) {
	query := squirrel.Update(tableName).
		SetMap(map[string]interface{}{
			"calendar":   user.CalendarId,
			"resume":     user.ResumeId,
			"name":       user.Name,
			"surname":    user.Surname,
			"patronymic": user.Patronymic,
			"email":      user.Email,
		}).
		Where(squirrel.Eq{"id": user.Id}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	res, err := query.ExecContext(ctx)
	if err != nil {
		return false, err
	}

	rows, err := res.RowsAffected()
	return rows != 0, err
}

func (r *repo) GetUsers(ctx context.Context, userIds []uint64) ([]models.User, error) {
	return nil, errNotImplemented
}
