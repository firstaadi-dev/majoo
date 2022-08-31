package usecase

import (
	"github.com/firstaadi-dev/majoo-backend-test/domain"
	"github.com/firstaadi-dev/majoo-backend-test/transaction"
)

type TransactionUsecase struct {
	TransactionRepository transaction.TransactionRepository
}

// MerchantByOutletID implements transaction.UseCase
func (t *TransactionUsecase) MerchantByOutletID(id int) (*domain.Merchant, error) {
	merchant, err := t.TransactionRepository.GetMerchantByOutletID(id)
	if err != nil {
		return nil, err
	}
	return merchant, nil
}

// ReportDailyMerchantOmzet implements transaction.UseCase
func (t *TransactionUsecase) ReportDailyMerchantOmzet(id int, page int) ([]domain.MerchantOmzet, error) {
	offset, limit := PageToOffsetLimit(page)
	report, err := t.TransactionRepository.GetMerchantOmzet(id, offset, limit)
	if err != nil {
		return nil, err
	}
	return report, nil
}

// ReportDailyOutletOmzet implements transaction.UseCase
func (t *TransactionUsecase) ReportDailyOutletOmzet(id int, date int) (*domain.OutletOmzet, error) {

	report, err := t.TransactionRepository.GetOutletOmzet(id, date)
	if err != nil {
		return nil, err
	}
	return report, nil

}

func PageToOffsetLimit(page int) (int, int) {
	var offset, limit int
	offset = (page - 1) * 10
	limit = 10
	return offset, limit
}

func NewTransactionUsecase(r transaction.TransactionRepository) transaction.UseCase {
	return &TransactionUsecase{
		TransactionRepository: r,
	}
}
