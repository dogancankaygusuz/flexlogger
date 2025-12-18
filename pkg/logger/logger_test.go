package logger

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestLoggerOutput(t *testing.T) {
	var buf bytes.Buffer

	logger := New(InfoLevel, &buf, &JSONFormatter{})
	logger.Info(context.Background(), "test message", nil)
	logger.Close()

	output := buf.String()

	if !strings.Contains(output, "test message") {
		t.Errorf("Log mesajı bulunamadı. Çıktı: %s", output)
	}
	if !strings.Contains(output, "INFO") {
		t.Errorf("Level yanlış. Çıktı: %s", output)
	}
}

func TestLogLevelFilter(t *testing.T) {
	var buf bytes.Buffer
	logger := New(InfoLevel, &buf, &TextFormatter{})

	logger.Debug(context.Background(), "gizli mesaj", nil)
	logger.Close()

	if buf.Len() > 0 {
		t.Errorf("Debug mesajı yazılmamalıydı ama yazıldı.")
	}
}
