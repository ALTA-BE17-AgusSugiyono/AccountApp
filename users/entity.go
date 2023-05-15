package users

import "time"

type Users struct {
	ID           int
	PhoneNumber  string
	Name         string
	Gender       string
	TanggalLahir string
	Balance      int
	DeletedAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
