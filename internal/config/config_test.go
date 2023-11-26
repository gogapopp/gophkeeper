package config

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	cfg, err := LoadConfig()
	assert.NoError(t, err)
	assert.IsType(t, &viper.Viper{}, cfg)
}
