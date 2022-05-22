package dao

import (
	"context"
	"database/sql"
	"fmt"
	"makretplace/internal/user/infrastructure/database/sqlc"
	"strings"

	"github.com/google/uuid"
)

type UserDao struct {
	*sqlc.Queries
	db *sql.DB
}

func NewUserDao(db *sql.DB) *UserDao {
	return &UserDao{
		db:      db,
		Queries: sqlc.New(db),
	}
}

type UpdateUserParams struct {
	Email    *string `db:"email"`
	Password *string `db:"password"`
	Name     *string `db:"name"`
}

func (u *UserDao) UpdateUser(ctx context.Context, userId uuid.UUID, arg UpdateUserParams) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if arg.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, *arg.Email)
		argId++
	}

	if arg.Password != nil {
		setValues = append(setValues, fmt.Sprintf("password=$%d", argId))
		args = append(args, *arg.Password)
		argId++
	}

	if arg.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *arg.Name)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE users SET %s WHERE id = $%d`,
		setQuery, argId)

	args = append(args, userId)

	_, err := u.db.ExecContext(ctx, query, args...)
	return err
}

const getUserInfoByID = `
SELECT id, name FROM "users" as "user" WHERE "user"."id" = $1
`

func (q *UserDao) GetUserInfoByID(ctx context.Context, id uuid.UUID) (sqlc.UserInfo, error) {
	row := q.db.QueryRowContext(ctx, getUserInfoByID, id)
	var i sqlc.UserInfo
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getUsersInfo = `
SELECT id, name FROM "users" as "user" LIMIT $1 OFFSET $2
`

type GetUsersInfoParams struct {
	Limit  int32 `db:"limit"`
	Offset int32 `db:"offset"`
}

func (q *UserDao) GetUsersInfo(ctx context.Context, arg GetUsersInfoParams) ([]sqlc.UserInfo, error) {
	rows, err := q.db.QueryContext(ctx, getUsersInfo, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sqlc.UserInfo
	for rows.Next() {
		var i sqlc.UserInfo
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
