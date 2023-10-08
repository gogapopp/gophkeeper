package hasher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
)

var ErrInvalidKey = errors.New("invalid hash key")

const sign = "secret_key+++"

// GeneratePasswordHas хеширует пароль
func GenerateHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(sign)))
}

func createKeyHash(key []byte) []byte {
	hasher := sha256.New()
	hasher.Write(key)
	return hasher.Sum(nil)
}

func Encrypt(plainText []byte, key []byte) ([]byte, error) {
	hashedKey := createKeyHash(key)
	block, err := aes.NewCipher(hashedKey)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plainText))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plainText)

	return ciphertext, nil
}

func Decrypt(cipherText []byte, key []byte) ([]byte, error) {
	hashedKey := createKeyHash(key)
	block, err := aes.NewCipher(hashedKey)
	if err != nil {
		return nil, err
	}

	if len(cipherText) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(cipherText, cipherText)
	return cipherText, nil
}
