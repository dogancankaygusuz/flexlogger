# FlexLogger 🚀

**FlexLogger**, Go uygulamaları için geliştirilmiş; yüksek performanslı, thread-safe, yapılandırılabilir ve üretim ortamına (production) uygun bir loglama kütüphanesidir.

Standart `log` paketinin ötesine geçerek, **Asenkron Yazma (Async)**, **Dosya Döndürme (Log Rotation)** ve **Context Takibi (Tracing)** gibi kurumsal özellikleri barındırır.

---

## 🌟 Özellikler

*   ⚡ **Asenkron & Non-Blocking:** `Channels` ve `Goroutines` kullanarak loglama işlemini arka planda yapar, ana akışı (latency) etkilemez.
*   🔄 **Otomatik Log Rotation:** Log dosyaları belirlenen boyuta (örn: 10MB) ulaştığında otomatik olarak yedeklenir (`app.log` -> `app-TIMESTAMP.backup`).
*   🔍 **Context Aware (Tracing):** `context.Context` desteği ile `request_id` veya `trace_id` gibi değerleri otomatik loglar.
*   🎨 **Çoklu Format Desteği:** 
    *   **JSON Formatter:** Log toplama araçları (ELK Stack, Splunk) için.
    *   **Text Formatter:** Geliştirme ortamı için renkli ve okunabilir çıktı.
*   🛡️ **Thread-Safe:** `sync.Mutex` ve `Worker Pattern` ile yüksek eşzamanlılık (concurrency) altında güvenle çalışır.
*   📍 **Caller Information:** Hatanın hangi dosya ve satırda olduğunu otomatik yakalar (örn: `main.go:42`).
*   📝 **Multi-Writer:** Logları aynı anda hem Dosyaya hem de Konsola yazabilir.

---

## 📦 Kurulum

```bash
go get github.com/dogancankaygusuz/flexlogger
