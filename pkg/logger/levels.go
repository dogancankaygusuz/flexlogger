package logger

import "strings"

// Level, log seviyesini temsil eden özel bir tiptir.
// int tabanlıdır çünkü seviyeleri karşılaştırmak isteyeceğiz (örn: level > InfoLevel)
type Level int

const (
	// iota ile otomatik artan değerler atıyoruz.
	// DebugLevel = 0, InfoLevel = 1, WarnLevel = 2 ...
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

// String metodu, Level tipini okunabilir metne çevirir.
// fmt.Stringer interface'ini implemente etmiş oluruz.
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

// ParseLevel, dışarıdan (config'den) gelen string'i Level tipine çevirir.
// "info" -> InfoLevel döner.
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
		return InfoLevel // Varsayılan olarak INFO olsun
	}
}
