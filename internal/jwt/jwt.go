package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const signingKEY = "secret_jwt_sign_key"

type JwtClaims struct {
	UserID int
	jwt.RegisteredClaims
}

func GenerateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	})
	return token.SignedString([]byte(signingKEY))
}

func ParseToken(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKEY), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims.UserID, nil
	} else {
		return 0, errors.New("invalid token")
	}
}
