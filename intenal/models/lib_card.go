package models

import "time"

type LibCard struct {
	ID        int64      `db:"id"`
	Code      string     `db:"code"`
	UserID    int64      `db:"user_id"`
	CreatedAt *time.Time `db:"created_at"`
	ExpiresAt *time.Time `db:"expires_at"`
}
