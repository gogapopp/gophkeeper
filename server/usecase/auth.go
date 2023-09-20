package usecase

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gogapopp/gophkeeper/internal/hash"
	"github.com/gogapopp/gophkeeper/models"
	"github.com/golang-jwt/jwt/v5"
)

const signingKEY = "secret_jwt_sign_key"

type (
	Auth interface {
		Register(ctx context.Context, user models.User) error
		Login(ctx context.Context, user models.User) (string, error)
	}

	JwtCLaims struct {
		UserID int
		jwt.RegisteredClaims
	}
)

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
	token, err := GenerateToken(userID)
	if err != nil {
		return "", fmt.Errorf("%s: %s", op, err)
	}
	return token, nil
}

func GenerateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCLaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	})
	return token.SignedString([]byte(signingKEY))
}
