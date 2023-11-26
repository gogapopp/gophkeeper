package file

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		assert.NoError(t, err)
	}
	defer os.Remove(tmpfile.Name())
	text := []byte("Hello, World!")
	if _, err := tmpfile.Write(text); err != nil {
		assert.NoError(t, err)
	}
	if err := tmpfile.Close(); err != nil {
		assert.NoError(t, err)
	}
	data, err := ReadFile(tmpfile.Name())
	assert.NoError(t, err)
	assert.Equal(t, text, data)
}
