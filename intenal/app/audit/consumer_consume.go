package audit

import (
	"context"
)

func (c *AuditConsumer) Consume(ctx context.Context) {

	if c.useBatch {
		c.ConsumeBatch(ctx)
		return
	}
	c.ConsumeSingle(ctx)
}
