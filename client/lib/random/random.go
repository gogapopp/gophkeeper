package random

import (
	"crypto/rand"
	"log"
	"strings"
)

func GenerateUniqueKey() string {
	const size = 8
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	var letters = []rune("0123456789")
	var result strings.Builder
	for _, b := range b {
		result.WriteRune(letters[int(b)%len(letters)])
	}
	return result.String()
}
