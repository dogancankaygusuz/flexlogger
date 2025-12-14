package logger

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type LogEntry struct {
	Level   Level                  `json:"level"`
	Time    time.Time              `json:"time"`
	Message string                 `json:"message"`
	Fields  map[string]interface{} `json:"fields,omitempty"`
}

type Formatter interface {
	Format(entry *LogEntry) ([]byte, error)
}

type JSONFormatter struct{}

func (f *JSONFormatter) Format(entry *LogEntry) ([]byte, error) {

	data := make(map[string]interface{})
	data["level"] = entry.Level.String()
	data["time"] = entry.Time.Format(time.RFC3339)
	data["message"] = entry.Message

	for k, v := range entry.Fields {
		data[k] = v
	}

	return json.Marshal(data)
}

type TextFormatter struct {
	UseColors bool
}

func (f *TextFormatter) Format(entry *LogEntry) ([]byte, error) {
	resetColor := "\033[0m"
	color := ""

	if f.UseColors {
		switch entry.Level {
		case DebugLevel:
			color = "\033[36m"
		case InfoLevel:
			color = "\033[32m"
		case WarnLevel:
			color = "\033[33m"
		case ErrorLevel, FatalLevel:
			color = "\033[31m"
		}
	}

	// Format: [ZAMAN] [LEVEL] Mesaj {field1=value1 field2=value2}
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	levelStr := entry.Level.String()

	// Fields kısmını string'e çevirme
	fieldsStr := ""
	if len(entry.Fields) > 0 {
		var sb strings.Builder
		for k, v := range entry.Fields {
			sb.WriteString(fmt.Sprintf("%s=%v ", k, v))
		}
		fieldsStr = fmt.Sprintf("\tFields: {%s}", strings.TrimSpace(sb.String()))
	}

	// Sonuç stringini oluştur
	logLine := fmt.Sprintf("%s[%s] [%s]%s %s%s\n",
		color, timestamp, levelStr, resetColor, entry.Message, fieldsStr)

	return []byte(logLine), nil
}
