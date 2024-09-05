package logger

import (
	"os"

	"github.com/Angstreminus/selector/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SelectLevel(cfg config.Config) zapcore.Level {
	switch cfg.LogLevel {
	case "Info":
		return zapcore.InfoLevel
	case "Debug":
		return zap.DebugLevel
	case "Error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func NewLogger(cfg config.Config) (*zap.Logger, error) {
	var logPlace zapcore.WriteSyncer
	if cfg.LogFilePath != "" {
		logPlace, err := os.OpenFile(cfg.LogFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
		if err != nil {
			return nil, err
		}
		defer logPlace.Close()
	} else {
		logPlace = os.Stdout
	}

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(SelectLevel(cfg))

	return zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			zapcore.Lock(logPlace),
			atomicLevel,
		),
	), nil
}
