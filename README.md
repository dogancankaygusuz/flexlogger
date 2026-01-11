# FlexLogger

**FlexLogger**, Go uygulamaları için geliştirilmiş; yüksek performanslı, thread-safe, yapılandırılabilir ve üretim ortamına uygun bir loglama kütüphanesidir.

Standart `log` paketinin ötesine geçerek, **Asenkron Yazma (Async)**, **Dosya Döndürme (Log Rotation)** ve **Context Takibi (Tracing)** gibi kurumsal özellikleri barındırır.

---

## Özellikler

*   **Asenkron & Non-Blocking:** `Channels` ve `Goroutines` kullanarak loglama işlemini arka planda yapar, ana akışı etkilemez.
*   **Otomatik Log Rotation:** Log dosyaları belirlenen boyuta (10MB) ulaştığında otomatik olarak yedeklenir (`app.log` -> `app-TIMESTAMP.backup`).
*   **Context Aware (Tracing):** `context.Context` desteği ile `request_id` veya `trace_id` gibi değerleri otomatik loglar.
*   **Thread-Safe:** `sync.Mutex` ve `Worker Pattern` ile yüksek eşzamanlılık (concurrency) altında güvenle çalışır.
*   **Caller Information:** Hatanın hangi dosya ve satırda olduğunu otomatik yakalar.
*   **Multi-Writer:** Logları aynı anda hem Dosyaya hem de Konsola yazabilir.

---

## Kurulum

```bash
go get github.com/dogancankaygusuz/flexlogger
```

## Kullanım

Geliştirme ortamında renkli konsol çıktısı için:

```go
package main

import (
	"context"
	"github.com/dogancankaygusuz/flexlogger/pkg/logger"
)

func main() {
	// Config ile Logger oluştur
	cfg := logger.Config{
		Level:      "DEBUG",
		Format:     logger.FormatText,
		UseConsole: true,
		UseColors:  true,
	}
	
	log, _ := logger.NewFromConfig(cfg)
	defer log.Close() // Uygulama kapanırken logları flush et

	// Loglama
	log.Info(context.Background(), "Uygulama başlatıldı", nil)
	log.Debug(context.Background(), "Bu bir debug mesajıdır", nil)
}
```

JSON formatı, dosya yazdırma, log rotation ve context takibi:

```go
package main

import (
	"context"
	"github.com/dogancankaygusuz/flexlogger/pkg/logger"
)

func main() {
	// Production Ayarları
	cfg := logger.Config{
		Level:      "INFO",
		Format:     logger.FormatJSON,
		FilePath:   "logs/app.log",
		UseFile:    true,  // Dosyaya yaz
		UseConsole: false, // Konsolu kapat
	}

	log, _ := logger.NewFromConfig(cfg)
	defer log.Close()

	// Context Simülasyonu (Request ID takibi)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "request_id", "req-xyz-123")

	// Structured Logging
	log.Info(ctx, "Sipariş alındı", map[string]interface{}{
		"user_id": 101,
		"amount":  59.90,
		"currency": "USD",
	})
	
	// Hata Logu (Otomatik dosya ve satır bilgisi eklenir)
	log.Error(ctx, "Ödeme başarısız", map[string]interface{}{
		"error_code": 5001,
	})
}
```

##  Mimari ve Tasarım Desenleri
Bu proje geliştirilirken aşağıdaki yazılım prensipleri ve tasarım desenleri kullanılmıştır:

* Strategy Pattern: JSONFormatter ve TextFormatter değişimleri için.
* Factory Pattern: NewFromConfig ile nesne oluşturma karmaşıklığını gizlemek için.
* Worker Pool Pattern: Logları asenkron işlemek için Goroutine ve Channel yapısı.
* Dependency Injection: io.Writer soyutlaması ile test edilebilir yapı.

## Yapılandırma

| Alan | Tip | Açıklama |
| :--- | :--- | :--- |
| `Level` | `string` | Log seviyesi (`DEBUG`, `INFO`, `WARN`, `ERROR`, `FATAL`) |
| `Format` | `FormatType` | Çıktı formatı (`logger.FormatJSON` veya `logger.FormatText`) |
| `FilePath` | `string` | Log dosyasının yolu |
| `UseConsole` | `bool` | Loglar konsola basılsın mı? |
| `UseFile` | `bool` | Loglar dosyaya kaydedilsin mi? |
| `UseColors` | `bool` | Text formatında renkli çıktı olsun mu? |

