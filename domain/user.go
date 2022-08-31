package domain

import "time"

type User struct {
	ID        int
	Name      string
	UserName  string
	Password  string
	CreatedAt time.Time
	CreatedBy int
	UpdatedAt time.Time
	UpdatedBy int
}
