package logger

import "strings"

type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

// String metodu, Level tipini okunabilir metne çevirir.
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// Config'den gelen string'i Level tipine çevirir.
func ParseLevel(levelStr string) Level {
	switch strings.ToUpper(levelStr) {
	case "DEBUG":
		return DebugLevel
	case "INFO":
		return InfoLevel
	case "WARN":
		return WarnLevel
	case "ERROR":
		return ErrorLevel
	case "FATAL":
		return FatalLevel
	default:
		return InfoLevel
	}
}
