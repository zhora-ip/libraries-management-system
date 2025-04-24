package audit

import (
	"sync"

	"github.com/zhora-ip/libraries-management-system/pkg"
)

// NewWP creates a new WorkerPool with specified writer and filtering options.
func NewWP(writer LogWriter, filtered bool) *WorkerPool {
	wp := &WorkerPool{
		jobs:       make(chan job, batchSize*2),
		numWorkers: numWorkers,
		workers:    make([]*worker, numWorkers),
		wg:         &sync.WaitGroup{},
	}

	var lw LogWriter
	if writer == nil {
		lw = newDefaultLogWriter()
	} else {
		lw = writer
	}

	for i := range numWorkers {
		wp.workers[i] = &worker{
			batch:     make([]job, 0, batchSize),
			batchSize: batchSize,
			writer:    lw,
		}

		if filtered {
			pkg.ParseConfig(&wp.workers[i].filter, configPath)
		}

	}
	return wp
}
