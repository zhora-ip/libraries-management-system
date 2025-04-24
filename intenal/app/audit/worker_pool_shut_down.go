package audit

import "log"

// ShutDown gracefully stops all workers and waits for completion.
func (p *WorkerPool) ShutDown() {
	if p.cancel != nil {
		p.cancel()
		p.wg.Wait()
		if p.next != nil {
			p.next.ShutDown()
		}
		log.Println("All workers exited gracefully")
	}
}
