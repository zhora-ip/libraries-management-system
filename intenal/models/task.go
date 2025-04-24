package models

import "time"

type TaskType int

const (
	TaskTypeUnknown  TaskType = iota // 0
	TaskTypeAuditLog                 // 1
)

type Task struct {
	ID           int
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	FinishedAt   *time.Time
	Status       TaskStatus
	AttemptCount int
	Type         TaskType
	Payload      []byte
}

type TaskStatus int

const (
	TaskStatusUnknown        TaskStatus = iota // 0
	TaskStatusCreated                          // 1
	TaskStatusProcessing                       // 2
	TaskStatusFailed                           // 3
	TaskStatusNoAttemptsLeft                   // 4
	Created                  = "CREATED"
	Processing               = "PROCESSING"
	Failed                   = "FAILED"
	NoAttemptsLeft           = "NO_ATTEMPTS_LEFT"
)

func (s TaskStatus) String() string {
	return [...]string{
		Unknown,
		Created,
		Processing,
		Failed,
		NoAttemptsLeft,
	}[s]
}

func TaskStatusFromString(status string) TaskStatus {
	switch status {
	case Created:
		return TaskStatusCreated
	case Processing:
		return TaskStatusProcessing
	case Failed:
		return TaskStatusFailed
	case NoAttemptsLeft:
		return TaskStatusNoAttemptsLeft
	default:
		return TaskStatusUnknown
	}
}
