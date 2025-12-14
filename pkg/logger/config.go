package logger

import (
	"errors"
	"os"
)

type OutputType string

const (
	OutputConsole OutputType = "CONSOLE"
	OutputFile    OutputType = "FILE"
)

type FormatType string

const (
	FormatJSON FormatType = "JSON"
	FormatText FormatType = "TEXT"
)

type Config struct {
	Level     string
	Output    OutputType
	Format    FormatType
	FilePath  string
	UseColors bool
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

	var writer *os.File
	switch cfg.Output {
	case OutputFile:
		if cfg.FilePath == "" {
			return nil, errors.New("output is FILE but FilePath is empty")
		}

		file, err := os.OpenFile(cfg.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
		writer = file
	case OutputConsole:
		fallthrough
	default:
		writer = os.Stdout
	}

	return New(level, writer, formatter), nil
}
