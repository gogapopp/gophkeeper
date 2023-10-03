package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/gogapopp/gophkeeper/models"
	"github.com/gogapopp/gophkeeper/server/repository"
)

func (r *Repository) Register(ctx context.Context, user models.User) error {
	const op = "postgres.auth.Register"
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (login, password, last_update_at) values ($1, $2, $3)", user.Login, user.Password, user.UploadedAt)
	if err != nil {
		if errors.As(err, &repository.PgErr) {
			switch repository.PgErr.Code {
			case "23505":
				return fmt.Errorf("%s: %s", op, repository.ErrUserAlreadyExists)
			default:
				return fmt.Errorf("%s: %s", op, err)
			}
		}
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

func (r *Repository) Login(ctx context.Context, user models.User) (string, error) {
	const op = "postgres.auth.Login"
	var userID int
	row := r.db.QueryRowContext(ctx, "SELECT user_id FROM users WHERE login=$1 AND password=$2", user.Login, user.Password)
	if err := row.Scan(&userID); err != nil {
		return "", fmt.Errorf("%s: %s", op, sql.ErrNoRows)
	}
	userIDstr := strconv.Itoa(userID)
	return userIDstr, nil
}
