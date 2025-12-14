package main

import (
	"os"

	"github.com/dogancankaygusuz/flexlogger/pkg/logger"
)

func main() {
	consoleLogger := logger.New(logger.DebugLevel, os.Stdout, &logger.TextFormatter{UseColors: true})

	consoleLogger.Info("Uygulama başlatıldı", nil)
	consoleLogger.Debug("Bu bir debug mesajıdır, detay içerir", map[string]interface{}{
		"version": "1.0.0",
		"env":     "dev",
	})

	jsonLogger := logger.New(logger.InfoLevel, os.Stdout, &logger.JSONFormatter{})

	jsonLogger.Info("JSON formatında log", map[string]interface{}{
		"status": 200,
		"module": "api",
	})

	jsonLogger.Debug("Bunu göremeyeceksin", nil)
}
