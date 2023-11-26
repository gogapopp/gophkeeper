package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateUniqueKey(t *testing.T) {
	keyMap := make(map[string]string)
	// заполняем мапу n кол-вом ключей, в цикле проверяем уникальность каждого
	for i := 0; i < 10; i++ {
		key := GenerateUniqueKey()
		assert.Equal(t, 8, len(key))
		// проверяем есть ли такой UniqueKey
		_, ok := keyMap[key]
		assert.False(t, ok)
		keyMap[key] = key
	}
}
