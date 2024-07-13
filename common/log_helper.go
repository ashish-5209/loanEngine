package common

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	CriticalLogger *zap.SugaredLogger
	Logger         *zap.SugaredLogger
	LogMap         map[string]string
)

func NewCriticalLogger(env, path string) *zap.SugaredLogger {
	var cfg zap.Config
	cfg.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	cfg.Encoding = "json"
	cfg.OutputPaths = []string{path}
	switch env {
	case "prod":
		cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "pp":
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "dev":
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	default:
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}
	logger, err := cfg.Build()
	if err != nil {
		log.Print("Failed to init critical logger", err)
	}
	return logger.Sugar()
}

func NewSugaredLogger(env, path string) *zap.SugaredLogger {
	var cfg zap.Config
	cfg.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	cfg.Encoding = "json"

	cfg.OutputPaths = []string{path}
	switch env {
	case "prod":
		cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "pp":
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "dev":
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	default:
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}
	logger, err := cfg.Build()
	if err != nil {
		log.Fatal("Failed to init logger", err)
	}
	return logger.Sugar()
}
