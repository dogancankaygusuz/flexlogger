package main

import (
	"context"
	"fmt"

	"github.com/dogancankaygusuz/flexlogger/pkg/logger"
)

func main() {
	// ==========================================
	// SENARYO 1: Production (Dosya + JSON)
	// ==========================================
	prodConfig := logger.Config{
		Level:      "INFO",
		Format:     logger.FormatJSON,
		FilePath:   "app.log",
		UseFile:    true,  // YENİ: Dosyaya yaz
		UseConsole: false, // Konsola yazma
	}

	prodLogger, err := logger.NewFromConfig(prodConfig)
	if err != nil {
		fmt.Println("Logger oluşturulamadı:", err)
		return
	}
	// Asenkron olduğu için program kapanmadan logları basması için:
	defer prodLogger.Close()

	// Context oluştur (request_id simülasyonu)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "request_id", "req-12345")

	prodLogger.Info(ctx, "Uygulama production modunda başladı", map[string]interface{}{
		"env": "production",
	})

	prodLogger.Error(ctx, "Veritabanı bağlantısı koptu", map[string]interface{}{
		"db_host": "192.168.1.50",
		"retry":   3,
	})

	fmt.Println("✅ Production logları 'app.log' dosyasına yazıldı.")

	// ==========================================
	// SENARYO 2: Development (Konsol + Renkli)
	// ==========================================
	devConfig := logger.Config{
		Level:      "DEBUG",
		Format:     logger.FormatText,
		UseColors:  true,
		UseConsole: true, // YENİ: Konsola yaz
		UseFile:    false,
	}

	devLogger, _ := logger.NewFromConfig(devConfig)
	defer devLogger.Close()

	devLogger.Debug(context.TODO(), "Bu bir debug mesajıdır", nil)

	devLogger.Warn(context.TODO(), "Disk alanı azalıyor", map[string]interface{}{
		"disk_usage": "85%",
	})
}
