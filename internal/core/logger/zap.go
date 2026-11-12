package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func MustNewZap() (*zap.Logger, func() error) {
	logger, close, err := NewZap()
	if err != nil {
		panic(fmt.Sprintf("failed to create zap logger: %v", err))
	}

	return logger, close
}

func NewZap() (*zap.Logger, func() error, error) {
	if err := os.MkdirAll("out/logs", 0755); err != nil {
		return nil, nil, fmt.Errorf("failed to create out/logs dir: %w", err)
	}

	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.000000")
	filepath := filepath.Join("out/logs", fmt.Sprintf("%s.log", timestamp))

	logfile, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create log file: %w", err)
	}

	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.000000")

	encoder := zapcore.NewConsoleEncoder(cfg)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.InfoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(logfile), zapcore.DebugLevel),
	)

	logger := zap.New(
		core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.WarnLevel),
	)

	return logger, logfile.Close, nil
}
