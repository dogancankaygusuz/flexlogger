package logger

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// LogEntry, bir log satırındaki tüm veriyi temsil eder.
type LogEntry struct {
	Level   Level                  `json:"level"`
	Time    time.Time              `json:"time"`
	Caller  string                 `json:"caller"` // YENİ: Dosya ve satır bilgisi (main.go:15)
	Message string                 `json:"message"`
	Fields  map[string]interface{} `json:"fields,omitempty"`
}

type Formatter interface {
	Format(entry *LogEntry) ([]byte, error)
}

// === JSON Formatter ===

type JSONFormatter struct{}

func (f *JSONFormatter) Format(entry *LogEntry) ([]byte, error) {
	data := make(map[string]interface{})
	data["level"] = entry.Level.String()
	data["time"] = entry.Time.Format(time.RFC3339)
	data["caller"] = entry.Caller // YENİ: JSON çıktısına ekle
	data["message"] = entry.Message

	for k, v := range entry.Fields {
		data[k] = v
	}

	return json.Marshal(data)
}

// === Text Formatter ===

type TextFormatter struct {
	UseColors bool
}

func (f *TextFormatter) Format(entry *LogEntry) ([]byte, error) {
	resetColor := "\033[0m"
	color := ""

	if f.UseColors {
		switch entry.Level {
		case DebugLevel:
			color = "\033[36m" // Cyan
		case InfoLevel:
			color = "\033[32m" // Green
		case WarnLevel:
			color = "\033[33m" // Yellow
		case ErrorLevel, FatalLevel:
			color = "\033[31m" // Red
		}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	levelStr := entry.Level.String()

	fieldsStr := ""
	if len(entry.Fields) > 0 {
		var sb strings.Builder
		for k, v := range entry.Fields {
			sb.WriteString(fmt.Sprintf("%s=%v ", k, v))
		}
		fieldsStr = fmt.Sprintf("\tFields: {%s}", strings.TrimSpace(sb.String()))
	}

	// YENİ FORMAT: [ZAMAN] [LEVEL] (DOSYA:SATIR) MESAJ
	// Caller bilgisini level'dan hemen sonraya ekledik.
	logLine := fmt.Sprintf("%s[%s] [%s] (%s)%s %s%s\n",
		color, timestamp, levelStr, entry.Caller, resetColor, entry.Message, fieldsStr)

	return []byte(logLine), nil
}
