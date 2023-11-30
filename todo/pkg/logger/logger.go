package logger

import (
	"sync"

	zap "go.uber.org/zap"
)

type upgradedLogger struct {
	*zap.SugaredLogger
}

var upgradedLoggerObject *upgradedLogger = nil
var singleOne sync.Once

func Get() *upgradedLogger {
	singleOne.Do(func() {
		baseLogger := createLoggerObject()
		upgradedLoggerObject = &upgradedLogger{
			baseLogger.Sugar(),
		}
	})
	return upgradedLoggerObject
}