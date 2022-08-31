package transaction

import "github.com/firstaadi-dev/majoo-backend-test/domain"

type TransactionRepository interface {
	GetMerchantOmzet(id, offset, limit int) ([]*domain.MerchantOmzet, error)
}
