package domain

import "time"

type Merchant struct {
	ID        int
	UserID    string
	Name      string
	CreatedAt time.Time
	CreatedBy int
	UpdatedAt time.Time
	UpdatedBy int
}
