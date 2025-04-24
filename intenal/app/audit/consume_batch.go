package audit

import (
	"context"
	"log"
)

func (c *AuditConsumer) ConsumeBatch(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			c.stop()
			return
		case <-c.consumer.GetDoneCh():
			c.stop()
			return
		default:
			msgs, err := c.consumer.PollBatch(timeoutConsumer)
			if err != nil {
				log.Print(err)
				continue
			}

			for _, msg := range msgs {
				c.wp.Submit(string(msg), nil)
			}
		}
	}

}
