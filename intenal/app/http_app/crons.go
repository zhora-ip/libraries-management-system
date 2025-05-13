package app

import (
	"context"
	"log"
	"time"
)

type oService interface {
	CheckCanceled(context.Context) error
	CheckExpired(context.Context) error
}

func findCanceledOrders(ctx context.Context, oService oService) {
	ticker := time.NewTicker(time.Second * timeoutCheckExpired)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := oService.CheckCanceled(ctx)
			if err != nil {
				log.Print(err)
			}
		}
	}
}

func findExpiredOrders(ctx context.Context, oService oService) {
	ticker := time.NewTicker(time.Second * timeoutCheckExpired)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := oService.CheckExpired(ctx)
			if err != nil {
				log.Print(err)
			}
		}
	}
}
