package logger

import (
	"io"
	"os"
)

type FormatType string

const (
	FormatJSON FormatType = "JSON"
	FormatText FormatType = "TEXT"
)

type Config struct {
	Level      string
	Format     FormatType
	FilePath   string
	UseConsole bool
	UseFile    bool
	UseColors  bool
}

func NewFromConfig(cfg Config) (*FlexLogger, error) {
	level := ParseLevel(cfg.Level)

	var formatter Formatter
	switch cfg.Format {
	case FormatJSON:
		formatter = &JSONFormatter{}
	case FormatText:
		formatter = &TextFormatter{UseColors: cfg.UseColors}
	default:
		formatter = &TextFormatter{UseColors: cfg.UseColors}
	}

	var writers []io.Writer

	if cfg.UseConsole {
		writers = append(writers, os.Stdout)
	}

	if cfg.UseFile && cfg.FilePath != "" {
		rotator := &SizeRotator{
			Filename: cfg.FilePath,
			MaxBytes: 10 * 1024 * 1024, // 10 MB
		}
		writers = append(writers, rotator)
	}

	finalOutput := io.Discard
	if len(writers) > 0 {
		finalOutput = io.MultiWriter(writers...)
	}

	return New(level, finalOutput, formatter), nil
}
