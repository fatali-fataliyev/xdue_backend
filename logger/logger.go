package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

var Logger *slog.Logger

func Init() error {
	currentDate := time.Now().Format("02_01_2006")
	logDir := "./logs"
	logFileName := currentDate + ".log"
	fullPath := filepath.Join(logDir, logFileName)

	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create %q directory: %w", logDir, err)
	}

	file, err := os.OpenFile(fullPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	multiWriter := io.MultiWriter(os.Stdout, file)

	opts := &slog.HandlerOptions{
		Level: slog.LevelError,
	}

	handler := slog.NewTextHandler(multiWriter, opts)

	Logger = slog.New(handler)

	slog.SetDefault(Logger)

	return nil
}
