package logger

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// LogEntry, bir log satırındaki tüm veriyi temsil eder.
// Formatter'lar bu yapıyı alıp işler.
type LogEntry struct {
	Level   Level                  `json:"level"`
	Time    time.Time              `json:"time"`
	Message string                 `json:"message"`
	Fields  map[string]interface{} `json:"fields,omitempty"` // Boşsa JSON'da gözükmesin
}

// Formatter, log verisini byte dizisine çeviren interface'dir.
type Formatter interface {
	Format(entry *LogEntry) ([]byte, error)
}

// ==========================================
// 1. JSON Formatter (Production için ideal)
// ==========================================

type JSONFormatter struct{}

func (f *JSONFormatter) Format(entry *LogEntry) ([]byte, error) {
	// LogEntry'yi doğrudan JSON'a çeviriyoruz.
	// Ancak Level int olduğu için (0,1,2) JSON'da sayı çıkar.
	// Bunu düzeltmek için anonim bir struct ile dönüşüm yapabiliriz.
	// Pratik olsun diye şimdilik struct'ı map'e çevirip basacağız.

	data := make(map[string]interface{})
	data["level"] = entry.Level.String() // "INFO", "ERROR" vs.
	data["time"] = entry.Time.Format(time.RFC3339)
	data["message"] = entry.Message

	// Varsayılan alanları ekle
	for k, v := range entry.Fields {
		data[k] = v
	}

	// JSON'a çevir (NewEncoder yerine Marshal kullanıyoruz şimdilik)
	return json.Marshal(data)
}

// ==========================================
// 2. Text Formatter (Local geliştirme için ideal)
// ==========================================

type TextFormatter struct {
	UseColors bool // Renkli çıktı istiyor muyuz?
}

func (f *TextFormatter) Format(entry *LogEntry) ([]byte, error) {
	// Renk kodları (ANSI escape codes)
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

	// Format: [ZAMAN] [LEVEL] Mesaj {field1=value1 field2=value2}
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	levelStr := entry.Level.String()

	// Fields kısmını string'e çevirelim (key=value formatında)
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
