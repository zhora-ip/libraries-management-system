package audit

// Submit sends a log for processing to the current pool and next pools.
// errCh is available only if a single worker pool without a pipelines is used,
// otherwise it must be nil.
func (p *WorkerPool) Submit(log any, errCh chan<- error) {

	p.jobs <- job{log, errCh}

	if p.next != nil {
		p.next.Submit(log, errCh)
	}

}

// SetNext sets the next WorkerPool in the pipeline.
func (p *WorkerPool) SetNext(w *WorkerPool) {
	p.next = w
}
