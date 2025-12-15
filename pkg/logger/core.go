package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath" // YENİ
	"runtime"       // YENİ
	"sync"
	"time"
)

// FlexLogger struct'ı ve New fonksiyonu AYNI KALIYOR...
// (Sadece log fonksiyonunu değiştireceğiz)

type FlexLogger struct {
	threshold Level
	output    io.Writer
	formatter Formatter
	mu        sync.Mutex
}

func New(threshold Level, output io.Writer, formatter Formatter) *FlexLogger {
	if output == nil {
		output = os.Stdout
	}
	if formatter == nil {
		formatter = &TextFormatter{UseColors: true}
	}

	return &FlexLogger{
		threshold: threshold,
		output:    output,
		formatter: formatter,
	}
}

// GÜNCELLENEN LOG FONKSİYONU
func (l *FlexLogger) log(level Level, msg string, fields map[string]interface{}) {
	if level < l.threshold {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	// YENİ: Caller (Çağıran) bilgisini al
	// skip=2 demek: runtime.Caller -> l.log -> l.Info -> KULLANICI KODU
	// Stack'te 2 adım yukarı çıkıyoruz.
	_, file, line, ok := runtime.Caller(2)
	caller := "unknown:0"
	if ok {
		// Full path yerine sadece dosya adını alalım (örn: /usr/go/src/main.go -> main.go)
		caller = fmt.Sprintf("%s:%d", filepath.Base(file), line)
	}

	entry := &LogEntry{
		Level:   level,
		Time:    time.Now(),
		Caller:  caller, // YENİ: Entry'e ekle
		Message: msg,
		Fields:  fields,
	}

	serialized, err := l.formatter.Format(entry)
	if err != nil {
		fmt.Printf("LOG FORMAT ERROR: %v\n", err)
		return
	}

	_, _ = l.output.Write(serialized)
}

// Interface methodları (Debug, Info vs.) AYNI KALIYOR...
func (l *FlexLogger) Debug(msg string, fields map[string]interface{}) { l.log(DebugLevel, msg, fields) }
func (l *FlexLogger) Info(msg string, fields map[string]interface{})  { l.log(InfoLevel, msg, fields) }
func (l *FlexLogger) Warn(msg string, fields map[string]interface{})  { l.log(WarnLevel, msg, fields) }
func (l *FlexLogger) Error(msg string, fields map[string]interface{}) { l.log(ErrorLevel, msg, fields) }
func (l *FlexLogger) Fatal(msg string, fields map[string]interface{}) {
	l.log(FatalLevel, msg, fields)
	os.Exit(1)
}
