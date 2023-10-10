package jwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParseToken(t *testing.T) {
	userID := 1
	token, err := GenerateToken(userID)
	assert.NoError(t, err)
	parsedUserID, err := ParseToken(token)
	assert.NoError(t, err)
	assert.Equal(t, userID, parsedUserID)
}
