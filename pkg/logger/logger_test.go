package logger

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestLoggerOutput(t *testing.T) {
	// Test için sanal bir buffer (Output) kullanacağız
	var buf bytes.Buffer

	// Logger oluştur
	logger := New(InfoLevel, &buf, &JSONFormatter{})

	// Log yaz
	logger.Info(context.Background(), "test message", nil)

	// Async olduğu için worker'ın yazmasını çok az beklemek gerekebilir
	// Ama test ortamında Close() çağırırsak flush eder.
	logger.Close()

	output := buf.String()

	// Kontroller
	if !strings.Contains(output, "test message") {
		t.Errorf("Log mesajı bulunamadı. Çıktı: %s", output)
	}
	if !strings.Contains(output, "INFO") {
		t.Errorf("Level yanlış. Çıktı: %s", output)
	}
}

func TestLogLevelFilter(t *testing.T) {
	var buf bytes.Buffer
	logger := New(InfoLevel, &buf, &TextFormatter{}) // Threshold INFO

	// DEBUG mesajı yaz (Yazılmamalı)
	logger.Debug(context.Background(), "gizli mesaj", nil)
	logger.Close()

	if buf.Len() > 0 {
		t.Errorf("Debug mesajı yazılmamalıydı ama yazıldı.")
	}
}
