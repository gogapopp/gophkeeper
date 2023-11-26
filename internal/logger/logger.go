package logger

import (
	"sync"

	"go.uber.org/zap"
)

var (
	errlogger error
	once      sync.Once
	logger    *zap.SugaredLogger
)

// SetupLogger устанавливает логгер
func SetupLogger() (*zap.SugaredLogger, error) {
	once.Do(func() {
		log, err := zap.NewProduction()
		if err != nil {
			errlogger = err
		}
		sugar := log.Sugar()
		logger = sugar
	})
	return logger, errlogger
}
