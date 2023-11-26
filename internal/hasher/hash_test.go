package hasher

import (
	"crypto/aes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	key := []byte("my secret key")
	plainText := []byte("Hello, World!")
	cipherText, err := Encrypt(plainText, key)
	assert.NoError(t, err)
	assert.NotEqual(t, plainText, cipherText[aes.BlockSize:])
	decyptedText, err := Decrypt(cipherText, key)
	assert.NoError(t, err)
	assert.Equal(t, plainText, decyptedText)
}
