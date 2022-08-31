package transaction

import "github.com/firstaadi-dev/majoo-backend-test/domain"

type UseCase interface {
	MerchantByOutletID(id int) (*domain.Merchant, error)
	ReportDailyMerchantOmzet(id, page int) ([]domain.MerchantOmzet, error)
	ReportDailyOutletOmzet(id, date int) (*domain.OutletOmzet, error)
}
