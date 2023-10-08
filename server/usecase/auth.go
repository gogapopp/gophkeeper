package usecase

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gogapopp/gophkeeper/internal/hasher"
	"github.com/gogapopp/gophkeeper/internal/jwt"
	"github.com/gogapopp/gophkeeper/models"
)

// Auth описывает взаимодействие с БД для аутентификации пользователя
type Auth interface {
	Register(ctx context.Context, user models.User) error
	Login(ctx context.Context, user models.User) (string, error)
}

// Register хешируем пароль и секретную фразу и возвращаем результат регистрации
func (au *AuthUsecase) Register(ctx context.Context, user models.User) error {
	const op = "usecase.auth.Register"
	user.Password = hasher.GenerateHash(user.Password)
	user.UserPhrase = hasher.GenerateHash(user.UserPhrase)
	err := au.auth.Register(ctx, user)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

// Login возвращает результат логина пользователя
func (au *AuthUsecase) Login(ctx context.Context, user models.User) (string, error) {
	const op = "usecase.auth.Login"
	user.Password = hasher.GenerateHash(user.Password)
	user.UserPhrase = hasher.GenerateHash(user.UserPhrase)
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
