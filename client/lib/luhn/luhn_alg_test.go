package luhn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckLuhn(t *testing.T) {
	assert.True(t, CheckLuhn("4532015112830366"))  // valid random visa number
	assert.False(t, CheckLuhn("1234567812345678")) // unvalid
}
