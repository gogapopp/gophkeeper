package usecase

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gogapopp/gophkeeper/internal/hash"
	"github.com/gogapopp/gophkeeper/internal/jwt"
	"github.com/gogapopp/gophkeeper/models"
)

type Auth interface {
	Register(ctx context.Context, user models.User) error
	Login(ctx context.Context, user models.User) (string, error)
}

func (au *AuthUsecase) Register(ctx context.Context, user models.User) error {
	const op = "usecase.auth.Register"
	user.Password = hash.GeneratePasswordHash(user.Password)
	err := au.auth.Register(ctx, user)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

func (au *AuthUsecase) Login(ctx context.Context, user models.User) (string, error) {
	const op = "usecase.auth.Login"
	user.Password = hash.GeneratePasswordHash(user.Password)
	user.UploadedAt = time.Now()
	userIDstr, err := au.auth.Login(ctx, user)
	if err != nil {
		return "", fmt.Errorf("%s: %s", op, err)
	}
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		return "", fmt.Errorf("%s: %s", op, err)
	}
	token, err := jwt.GenerateToken(userID)
	if err != nil {
		return "", fmt.Errorf("%s: %s", op, err)
	}
	return token, nil
}