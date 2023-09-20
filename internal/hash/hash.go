package hash

import (
	"crypto/sha256"
	"fmt"
)

const sign = "secret_key+++"

// GeneratePasswordHas хеширует пароль
func GeneratePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(sign)))
}
