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

type MerchantOmzet struct {
	Name  string
	Omzet int
}

type OutletOmzet struct {
	MerchantName string
	OutletName   string
	Omzet        int
}
