package logger

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type SizeRotator struct {
	Filename    string
	MaxBytes    int64
	currentFile *os.File
	currentSize int64
	mu          sync.Mutex
}

func (r *SizeRotator) Write(p []byte) (n int, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.currentFile == nil {
		if err := r.openFile(); err != nil {
			return 0, err
		}
	}

	writeLen := int64(len(p))
	if r.currentSize+writeLen > r.MaxBytes {
		if err := r.rotate(); err != nil {
			return 0, err
		}
	}

	n, err = r.currentFile.Write(p)
	r.currentSize += int64(n)
	return n, err
}

func (r *SizeRotator) openFile() error {
	file, err := os.OpenFile(r.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	info, err := file.Stat()
	if err != nil {
		file.Close()
		return err
	}
	r.currentFile = file
	r.currentSize = info.Size()
	return nil
}

func (r *SizeRotator) rotate() error {
	if err := r.currentFile.Close(); err != nil {
		return err
	}
	timestamp := time.Now().Format("2006-01-02T15-04-05")
	newName := fmt.Sprintf("%s-%s.backup", r.Filename, timestamp)
	if err := os.Rename(r.Filename, newName); err != nil {
		return err
	}
	return r.openFile()
}

func (r *SizeRotator) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.currentFile != nil {
		return r.currentFile.Close()
	}
	return nil
}
