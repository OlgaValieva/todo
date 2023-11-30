package logger

import (
	"os"

	zap "go.uber.org/zap"

	"go.uber.org/zap/zapcore"
)

func createLoggerObject() *zap.Logger {
	hightPriority :=zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority :=zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	consoleEncoderConfig := zapcore.EncoderConfig{
		MessageKey: "message",

		LevelKey: "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,

		TimeKey: "time",
		EncodeTime: zapcore.ISO8601TimeEncoder,

		CallerKey: "caller",
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	jsonEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)

	stdOutput:= zapcore.AddSync(os.Stdout)

	core := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder,stdOutput, lowPriority),
		zapcore.NewCore(consoleEncoder,stdOutput, hightPriority),
	)
	loggerObject := zap.New(core)
	return loggerObject
}