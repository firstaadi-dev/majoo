package transaction

import "github.com/firstaadi-dev/majoo-backend-test/domain"

type TransactionRepository interface {
	GetMerchantByOutletID(id int) (*domain.Merchant, error)
	// GetMerchantOmzet(id, offset, limit int) ([]domain.MerchantOmzet, error)
	GetMerchantOmzet(id, date int) (*domain.MerchantOmzet, error)
	GetOutletOmzet(id, date int) (*domain.OutletOmzet, error)
}
