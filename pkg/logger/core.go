package logger

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

type FlexLogger struct {
	threshold Level
	output    io.Writer
	formatter Formatter
	logChan   chan *LogEntry
	wg        sync.WaitGroup
}

// Constructor
func New(threshold Level, output io.Writer, formatter Formatter) *FlexLogger {
	if output == nil {
		output = os.Stdout
	}
	if formatter == nil {
		formatter = &TextFormatter{UseColors: true}
	}

	l := &FlexLogger{
		threshold: threshold,
		output:    output,
		formatter: formatter,
		logChan:   make(chan *LogEntry, 1000), // 1000 log kapasiteli buffer
	}

	l.wg.Add(1)
	go l.startWorker()

	return l
}

func (l *FlexLogger) startWorker() {
	defer l.wg.Done()

	for entry := range l.logChan {
		serialized, err := l.formatter.Format(entry)
		if err == nil {
			_, _ = l.output.Write(serialized)
		}
	}
}

func (l *FlexLogger) Close() {
	close(l.logChan)
	l.wg.Wait()
}

func (l *FlexLogger) log(ctx context.Context, level Level, msg string, fields map[string]interface{}) {
	if level < l.threshold {
		return
	}

	if ctx != nil {
		if traceID, ok := ctx.Value("request_id").(string); ok {
			if fields == nil {
				fields = make(map[string]interface{})
			}
			fields["request_id"] = traceID
		}
	}

	_, file, line, ok := runtime.Caller(2)
	caller := "unknown:0"
	if ok {
		caller = fmt.Sprintf("%s:%d", filepath.Base(file), line)
	}

	entry := &LogEntry{
		Level:   level,
		Time:    time.Now(),
		Caller:  caller,
		Message: msg,
		Fields:  fields,
	}

	select {
	case l.logChan <- entry:
	default:
		fmt.Printf("LOG DROP: Channel full. Message: %s\n", msg)
	}
}

// Interface metotları
func (l *FlexLogger) Debug(ctx context.Context, msg string, fields map[string]interface{}) {
	l.log(ctx, DebugLevel, msg, fields)
}
func (l *FlexLogger) Info(ctx context.Context, msg string, fields map[string]interface{}) {
	l.log(ctx, InfoLevel, msg, fields)
}
func (l *FlexLogger) Warn(ctx context.Context, msg string, fields map[string]interface{}) {
	l.log(ctx, WarnLevel, msg, fields)
}
func (l *FlexLogger) Error(ctx context.Context, msg string, fields map[string]interface{}) {
	l.log(ctx, ErrorLevel, msg, fields)
}
func (l *FlexLogger) Fatal(ctx context.Context, msg string, fields map[string]interface{}) {
	l.log(ctx, FatalLevel, msg, fields)
	l.Close()
	os.Exit(1)
}
