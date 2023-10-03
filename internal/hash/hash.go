package hash

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
)

const sign = "secret_key+++"

// GeneratePasswordHas хеширует пароль
func GeneratePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(sign)))
}

func Encrypt(plainText []byte, key []byte) ([]byte, error) {
	hash := sha256.Sum256(key)
	block, err := aes.NewCipher(hash[:])
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plainText)+sha256.Size)
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plainText)

	hash = sha256.Sum256(plainText)
	copy(ciphertext[aes.BlockSize+len(plainText):], hash[:])

	return ciphertext, nil
}

func Decrypt(cipherText []byte, key []byte) ([]byte, error) {
	hash := sha256.Sum256(key)
	block, err := aes.NewCipher(hash[:])
	if err != nil {
		return nil, err
	}

	if len(cipherText) < aes.BlockSize+sha256.Size {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	hash = sha256.Sum256(cipherText[:len(cipherText)-sha256.Size])
	if string(hash[:]) != string(cipherText[len(cipherText)-sha256.Size:]) {
		return nil, fmt.Errorf("invalid hash")
	}

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(cipherText, cipherText)
	return cipherText[:len(cipherText)-sha256.Size], nil
}
