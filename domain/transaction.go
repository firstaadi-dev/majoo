package domain

import "time"

type Transaction struct {
	ID         int
	MerchantID int
	OutletID   int
	BillTotal  int
	CreatedAt  time.Time
	CreatedBy  int
	UpdatedAt  time.Time
	UpdatedBy  int
}
 