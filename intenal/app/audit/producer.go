package audit

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

const (
	timeoutFindLogs = 1
	timeoutFailed   = 2
	maxAttempts     = 3
)

type producer interface {
	Produce([]byte) error
}

type tasksRepo interface {
	FindTasks(context.Context) ([]*models.Task, error)
	Update(context.Context, models.Task) error
	Delete(context.Context, int) error
}

type AuditProducer struct {
	tasksRepo tasksRepo
	producer  producer
}

func NewAuditProducer(tasksRepo tasksRepo, producer producer) *AuditProducer {
	return &AuditProducer{
		tasksRepo: tasksRepo,
		producer:  producer,
	}
}
