package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestLogger(t *testing.T) {
	logger, err := SetupLogger()
	assert.NoError(t, err)
	assert.IsType(t, &zap.SugaredLogger{}, logger)
}
