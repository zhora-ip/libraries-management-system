package audit

import "context"

// Run starts all workers in the pool and subsequent pools.
func (p *WorkerPool) Run() {
	p.ctx, p.cancel = context.WithCancel(context.Background())

	if p.next != nil {
		p.next.Run()
	}

	for i, w := range p.workers {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			w.work(p.ctx, i+1, p.jobs)
		}()

	}
}
