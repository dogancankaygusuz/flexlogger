package main

import (
	"fmt"

	"github.com/dogancankaygusuz/flexlogger/pkg/logger"
)

func main() {
	prodConfig := logger.Config{
		Level:    "INFO",
		Output:   logger.OutputFile,
		Format:   logger.FormatJSON,
		FilePath: "app.log",
	}

	prodLogger, err := logger.NewFromConfig(prodConfig)
	if err != nil {
		fmt.Println("Logger oluşturulamadı:", err)
		return
	}

	prodLogger.Info("Uygulama production modunda başladı", map[string]interface{}{
		"env": "production",
	})

	prodLogger.Error("Veritabanı bağlantısı koptu", map[string]interface{}{
		"db_host": "192.168.1.50",
		"retry":   3,
	})

	fmt.Println("✅ Production logları 'app.log' dosyasına yazıldı.")

	devConfig := logger.Config{
		Level:     "DEBUG",
		Output:    logger.OutputConsole,
		Format:    logger.FormatText,
		UseColors: true,
	}

	devLogger, _ := logger.NewFromConfig(devConfig)

	devLogger.Debug("Bu bir debug mesajıdır", nil)
	devLogger.Warn("Disk alanı azalıyor", map[string]interface{}{
		"disk_usage": "85%",
	})
}
