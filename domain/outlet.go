package domain

import "time"

type Outlet struct {
	ID         int
	MerchantID int
	Name       string
	CreatedAt  time.Time
	CreatedBy  int
	UpdatedAt  time.Time
	UpdatedBy  int
}
