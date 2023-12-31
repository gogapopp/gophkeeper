package postgres

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/gogapopp/gophkeeper/models"
	"github.com/gogapopp/gophkeeper/server/repository"
	"github.com/jackc/pgerrcode"
)

// Register сохраняем данные пользователя
func (r *Repository) Register(ctx context.Context, user models.User) error {
	const op = "postgres.auth.Register"
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (login, password, user_phrase, last_update_at) values ($1, $2, $3, $4)", user.Login, user.Password, user.UserPhrase, user.UploadedAt)
	if err != nil {
		if errors.As(err, &repository.PgErr) {
			switch repository.PgErr.Code {
			case pgerrcode.UniqueViolation:
				return fmt.Errorf("%s: %w", op, repository.ErrUserAlreadyExists)
			default:
				return fmt.Errorf("%s: %w", op, err)
			}
		}
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

// Login получает айди пользователя по паролю и логину
func (r *Repository) Login(ctx context.Context, user models.User) (string, error) {
	const op = "postgres.auth.Login"
	var userID int
	row := r.db.QueryRowContext(ctx, "SELECT user_id FROM users WHERE login=$1 AND password=$2 AND user_phrase=$3", user.Login, user.Password, user.UserPhrase)
	if err := row.Scan(&userID); err != nil {
		return "", fmt.Errorf("%s: %w", op, repository.ErrUserNotExists)
	}
	userIDstr := strconv.Itoa(userID)
	return userIDstr, nil
}
