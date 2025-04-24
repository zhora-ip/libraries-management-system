package audit

import (
	"context"
	"log"
	"os"
)

// LogWriter defines the interface for audit log writing.
type LogWriter interface {
	// Add writes a log to storage and returns count of written logs and error.
	Add(ctx context.Context, log string) (int, error)
}

type defaultLogWriter struct {
	log.Logger
}

func newDefaultLogWriter() *defaultLogWriter {
	return &defaultLogWriter{*log.New(os.Stdout, "", 0)}
}

// Add implements logWriter by printing logs to stdout.
func (l *defaultLogWriter) Add(_ context.Context, log string) (int, error) {
	l.Print(log)
	return 0, nil
}
