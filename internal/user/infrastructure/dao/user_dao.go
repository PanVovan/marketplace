package dao

import (
	"context"
	"database/sql"
	"fmt"
	"main/internal/user/infrastructure/sqlc"
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
