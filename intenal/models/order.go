package models

import "time"

type Status int

const (
	StatusUnknown   Status = iota // 0
	StatusAvailable               // 1
	StatusExpired                 // 2
	StatusIssued                  // 3
	Unknown         = "UNKNOWN"
	Available       = "AVAILABLE"
	Expired         = "EXPIRED"
	Issued          = "ISSUED"
)

func (s Status) String() string {
	return [...]string{
		Unknown,
		Available,
		Expired,
		Issued,
	}[s]
}

type Order struct {
	ID             int64      `db:"id"`
	LibraryID      int64      `db:"library_id"`
	PhysicalBookID int64      `db:"physical_book_id"`
	UserID         int64      `db:"user_id"`
	Status         Status     `db:"status"`
	CreatedAt      *time.Time `db:"created_at"`
	UpdatedAt      *time.Time `db:"updated_at"`
	ExpiresAt      *time.Time `db:"expires_at"`
}
