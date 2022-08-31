package usecase

import (
	"github.com/firstaadi-dev/majoo-backend-test/domain"
	"github.com/firstaadi-dev/majoo-backend-test/transaction"
)

type TransactionUsecase struct {
	TransactionRepository transaction.TransactionRepository
}

// ReportDailyMerchantOmzet implements transaction.UseCase
func (t *TransactionUsecase) ReportDailyMerchantOmzet(id int, page int) ([]*domain.MerchantOmzet, error) {
	// user, err := helper.ParseToken(accessToken)
	// if err != nil {
	// 	return nil, err
	// }
	// if id != user.ID {
	// 	return nil, fmt.Errorf("can't access another merchant report")
	// }
	var offset, limit int
	offset = page * 9
	limit = 10
	report, err := t.TransactionRepository.GetMerchantOmzet(id, offset, limit)
	if err != nil {
		return nil, err
	}
	return report, nil
}

func NewTransactionUsecase(r transaction.TransactionRepository) transaction.UseCase {
	return &TransactionUsecase{
		TransactionRepository: r,
	}
}
