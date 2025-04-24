package audit

import (
	"context"
	"log"
)

func (c *AuditConsumer) ConsumeSingle(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			c.stop()
			return
		case <-c.consumer.GetDoneCh():
			c.stop()
			return
		default:
			msg, err := c.consumer.Poll(timeoutConsumer)
			if err != nil {
				log.Print(err)
				return
			}

			if msg == nil {
				continue
			}

			c.wp.Submit(string(msg), nil)

		}
	}
}

func (c *AuditConsumer) stop() {
	log.Print("Shutting down consumer loop...")
	stopCh := c.consumer.GetStopCh()
	stopCh <- struct{}{}
}
