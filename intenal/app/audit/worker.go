package audit

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

type worker struct {
	batch     []job
	batchSize int
	writer    LogWriter
	filter    []string
}

func (w *worker) work(ctx context.Context, id int, jobs <-chan job) {

	ticker := time.NewTicker(time.Millisecond * timeout)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			w.process(ctx)
			log.Printf("Worker %d exited gracefully", id)
			return

		case val, ok := <-jobs:
			if !ok {
				w.process(ctx)
				return
			}

			w.batch = append(w.batch, val)

			if len(w.batch) >= w.batchSize {
				w.process(ctx)
			}
			ticker.Reset(time.Millisecond * timeout)

		case <-ticker.C:
			w.process(ctx)
			ticker.Reset(time.Millisecond * timeout)

		}
	}
}

func (w *worker) process(ctx context.Context) {

	for _, v := range w.batch {

		log := ""

		switch val := v.job.(type) {
		case *models.AuditStatusChange:
			log = val.String()
		case *models.AuditRequest:
			log = val.String()
		case *models.AuditResponse:
			log = val.String()
		case string:
			log = val
		default:
			continue
		}

		if w.filter != nil && containsAny(log, w.filter) || w.filter == nil {
			_, err := w.writer.Add(ctx, log)
			if v.errCh != nil {
				v.errCh <- err
			}

		}
	}
	w.batch = w.batch[:0]
}

func containsAny(s string, substrs []string) bool {
	for _, substr := range substrs {
		if strings.Contains(strings.ToLower(s), strings.ToLower(substr)) {
			return true
		}
	}
	return false
}
