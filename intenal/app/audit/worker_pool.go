package audit

import (
	"context"
	"sync"
)

const (
	batchSize  = 5
	numWorkers = 2
	timeout    = 500
)

var (
	configPath = "configs/filter.yaml"
)

type job struct {
	job   any
	errCh chan<- error
}

// WorkerPool represents a pool of workers processing audit logs.
type WorkerPool struct {
	ctx        context.Context
	cancel     context.CancelFunc
	wg         *sync.WaitGroup
	jobs       chan job
	next       *WorkerPool
	workers    []*worker
	numWorkers int
}
