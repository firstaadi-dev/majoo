package transaction

import "github.com/firstaadi-dev/majoo-backend-test/domain"

type UseCase interface {
	ReportDailyMerchantOmzet(id, page int) ([]*domain.MerchantOmzet, error)
	// ReportDailyOutletOmzet(id, page int)
}
