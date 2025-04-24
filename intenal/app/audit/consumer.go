package audit

const (
	timeoutConsumer = 100
)

type consumer interface {
	Poll(int) ([]byte, error)
	PollBatch(int) ([][]byte, error)
	GetDoneCh() <-chan struct{}
	GetStopCh() chan<- struct{}
}

type workerPool interface {
	Submit(any, chan<- error)
}

type AuditConsumer struct {
	consumer consumer
	wp       workerPool
	useBatch bool
}

func NewAuditConsumer(consumer consumer, wp workerPool, useBatch bool) *AuditConsumer {
	return &AuditConsumer{
		consumer: consumer,
		wp:       wp,
		useBatch: useBatch,
	}
}
