package random

import (
	"crypto/rand"
	"strings"
)

// GenerateUniqueKey генерирует уникальным ключ состоящий из 8 цифр
func GenerateUniqueKey() string {
	const size = 8
	b := make([]byte, size)
	_, _ = rand.Read(b)
	var letters = []rune("0123456789")
	var result strings.Builder
	for _, b := range b {
		result.WriteRune(letters[int(b)%len(letters)])
	}
	return result.String()
}
