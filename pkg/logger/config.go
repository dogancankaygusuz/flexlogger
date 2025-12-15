package logger

import (
	"io"
	"os"
)

// ==========================================
// EKSİK OLAN KISIMLAR EKLENDİ (BURASI)
// ==========================================

// FormatType, logun formatını belirler.
type FormatType string

const (
	FormatJSON FormatType = "JSON"
	FormatText FormatType = "TEXT"
)

// ==========================================

// Config yapısı
type Config struct {
	Level      string
	Format     FormatType // Yukarıda tanımladığımız tipi kullanıyoruz
	FilePath   string
	UseConsole bool // Konsola yazsın mı?
	UseFile    bool // Dosyaya yazsın mı?
	UseColors  bool // Renkli çıktı olsun mu?
}

// NewFromConfig Factory Method
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

	// Writers listesi oluştur
	var writers []io.Writer

	// 1. Konsol isteniyorsa ekle
	if cfg.UseConsole {
		writers = append(writers, os.Stdout)
	}

	// 2. Dosya isteniyorsa ekle
	if cfg.UseFile && cfg.FilePath != "" {
		// Log Rotation özelliğini kullanıyoruz (rotator.go dosyasındaki struct)
		rotator := &SizeRotator{
			Filename: cfg.FilePath,
			MaxBytes: 10 * 1024 * 1024, // 10 MB
		}
		writers = append(writers, rotator)
	}

	// MultiWriter ile birleştir
	// Eğer hiç writer yoksa (writers boşsa) io.Discard (hiçbir yere yazma) kullanılır.
	finalOutput := io.Discard
	if len(writers) > 0 {
		finalOutput = io.MultiWriter(writers...)
	}

	return New(level, finalOutput, formatter), nil
}
