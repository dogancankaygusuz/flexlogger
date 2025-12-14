package logger

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

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

func (l *FlexLogger) log(level Level, msg string, fields map[string]interface{}) {

	if level < l.threshold {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	entry := &LogEntry{
		Level:   level,
		Time:    time.Now(),
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

func (l *FlexLogger) Debug(msg string, fields map[string]interface{}) {
	l.log(DebugLevel, msg, fields)
}

func (l *FlexLogger) Info(msg string, fields map[string]interface{}) {
	l.log(InfoLevel, msg, fields)
}

func (l *FlexLogger) Warn(msg string, fields map[string]interface{}) {
	l.log(WarnLevel, msg, fields)
}

func (l *FlexLogger) Error(msg string, fields map[string]interface{}) {
	l.log(ErrorLevel, msg, fields)
}

func (l *FlexLogger) Fatal(msg string, fields map[string]interface{}) {
	l.log(FatalLevel, msg, fields)
	os.Exit(1)
}
