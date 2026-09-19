package core_logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger

	file *os.File
}

func NewLogger(config Config) (*Logger, error) {
	zapLvl := zap.NewAtomicLevel()
	if err := zapLvl.UnmarshalText([]byte(config.Level)); err != nil {
		return nil, fmt.Errorf("Log Level: %w", err)
	}

	if err := os.MkdirAll(config.Folder, 0755); err != nil {
		return nil, fmt.Errorf("Log Level: %w", err)
	}
	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.000000")
	LogFilePath := filepath.Join(
		config.Folder,
		fmt.Sprintf("%s.Log", timestamp),
	)

	LogFile, err := os.OpenFile(LogFilePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Log Level: %w", err)
	}

	zapConfig := zap.NewDevelopmentEncoderConfig()
	zapConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.000000")
	zapEncoder := zapcore.NewConsoleEncoder(zapConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(zapEncoder, zapcore.AddSync(os.Stdout), zapLvl),
		zapcore.NewCore(zapEncoder, zapcore.AddSync(LogFile), zapLvl),
	)

	zapLogger := zap.New(core, zap.AddCaller())

	return &Logger{
		Logger: zapLogger,
		file:   LogFile,
	}, nil
}

func (l *Logger) Close() {
	if err := l.file.Close(); err != nil {
		fmt.Println("failed to close logger: ", err)
	}
}
