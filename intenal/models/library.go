package models

import "time"

type Library struct {
	ID          int64      `db:"id"`
	Name        string     `db:"name"`
	Address     string     `db:"address"`
	PhoneNumber string     `db:"phone_number"`
	Lat         float64    `db:"lat"`
	Lng         float64    `db:"lng"`
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}
