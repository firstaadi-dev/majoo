package mysql

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/firstaadi-dev/majoo-backend-test/domain"
	"github.com/firstaadi-dev/majoo-backend-test/transaction"
)

type mysqlTransactionRepo struct {
	DB *sql.DB
}

// GetMerchantByOutletID implements transaction.TransactionRepository
func (t *mysqlTransactionRepo) GetMerchantByOutletID(id int) (*domain.Merchant, error) {
	var m domain.Merchant
	err := t.DB.QueryRow("SELECT merchant_id FROM `Outlets` WHERE id = ?", id).Scan(&m.ID)
	if err != nil {
		return nil, fmt.Errorf("outlet not found")
	}
	return &m, nil
}

// GetOutletOmzet implements transaction.TransactionRepository
func (t *mysqlTransactionRepo) GetOutletOmzet(id int, date int) (*domain.OutletOmzet, error) {
	var report domain.OutletOmzet
	var dateStart = "2021-11-" + strconv.Itoa(date) + " 00:00:00"
	var dateEnd = "2021-11-" + strconv.Itoa(date) + " 23:59:59"

	err := t.DB.QueryRow("SELECT "+
		"m.merchant_name AS merchant_name, "+
		"MAX(o.outlet_name) AS outlet_name, "+
		"SUM(t.bill_total) AS omzet "+
		"FROM `Transactions` AS t "+
		"INNER JOIN `Merchants` AS m ON t.merchant_id = m.id "+
		"INNER JOIN `Outlets` AS o ON t.outlet_id = o.id "+
		"WHERE "+
		"t.created_at BETWEEN ? AND ? "+
		"AND t.outlet_id = ? "+
		"GROUP BY merchant_name LIMIT 1", dateStart, dateEnd, id).Scan(
		&report.MerchantName,
		&report.OutletName,
		&report.Omzet,
	)

	if err == sql.ErrNoRows {
		err := t.DB.QueryRow("SELECT "+
			"m.merchant_name AS merchant_name, "+
			"o.outlet_name AS outlet_name "+
			"FROM `Transactions` AS t "+
			"INNER JOIN `Merchants` AS m ON t.merchant_id = m.id "+
			"INNER JOIN `Outlets` AS o ON t.outlet_id = o.id "+
			"WHERE t.outlet_id = ? "+
			"LIMIT 1", id).Scan(
			&report.MerchantName,
			&report.OutletName,
		)
		if err != nil {
			return nil, err
		}
		report = domain.OutletOmzet{
			MerchantName: report.MerchantName,
			OutletName:   report.OutletName,
			Omzet:        0,
		}
	}

	return &report, nil
}

// GetMerchantOmzet implements transaction.TransactionRepository
func (t *mysqlTransactionRepo) GetMerchantOmzet(id int, date int) (*domain.MerchantOmzet, error) {
	var report domain.MerchantOmzet
	var dateStart = "2021-11-" + strconv.Itoa(date) + " 00:00:00"
	var dateEnd = "2021-11-" + strconv.Itoa(date) + " 23:59:59"
	err := t.DB.QueryRow("SELECT "+
		"m.merchant_name, "+
		// "DAY(t.created_at) as transaction_date, "+
		"SUM(t.bill_total) AS omzet "+
		"FROM `Transactions` AS t "+
		"INNER JOIN `Merchants` as m ON t.merchant_id = m.id "+
		"WHERE "+
		"t.created_at BETWEEN ? AND ? "+
		"AND t.merchant_id = ? "+
		"GROUP BY "+
		"m.merchant_name", dateStart, dateEnd, id).Scan(
		&report.Name,
		&report.Omzet,
	)

	if err == sql.ErrNoRows {
		err := t.DB.QueryRow("SELECT merchant_name FROM `Merchants` WHERE id = ?", id).Scan(&report.Name)
		if err != nil {
			return nil, err
		}
		report.Omzet = 0
	}

	return &report, nil
}

// // GetMerchantOmzet implements transaction.TransactionRepository
// func (t *mysqlTransactionRepo) GetMerchantOmzet(id int, offset int, limit int) ([]domain.MerchantOmzet, error) {
// 	var reportMap = make(map[int]domain.MerchantOmzet)
// 	type merchant struct {
// 		Name string
// 	}
// 	var m merchant
// 	err := t.DB.QueryRow("SELECT merchant_name FROM `Merchants` WHERE id = ?", id).Scan(&m.Name)
// 	if err != nil {
// 		return nil, err
// 	}
// 	for i := 1; i <= 30; i++ {
// 		reportMap[i] = domain.MerchantOmzet{
// 			Name:  m.Name,
// 			Day:   i,
// 			Omzet: 0,
// 		}
// 	}
// 	var reportArr = make([]domain.MerchantOmzet, 0, 10)
// 	results, err := t.DB.Query("SELECT "+
// 		"m.merchant_name, "+
// 		"DAY(t.created_at) as transaction_date, "+
// 		"SUM(t.bill_total) AS omzet "+
// 		"FROM `Transactions` AS t "+
// 		"INNER JOIN `Merchants` as m ON t.merchant_id = m.id "+
// 		"WHERE "+
// 		"t.created_at BETWEEN '2021-11-01 00:00:00' AND '2021-11-30 23:59:59' "+
// 		"AND t.merchant_id = ? "+
// 		"GROUP BY "+
// 		"transaction_date, "+
// 		"m.merchant_name "+
// 		"ORDER BY transaction_date ", id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	for results.Next() {
// 		var report domain.MerchantOmzet
// 		err = results.Scan(
// 			&report.Name,
// 			&report.Day,
// 			&report.Omzet,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}
// 		reportMap[report.Day] = report
// 	}
// 	if offset+limit >= 30 {
// 		limit = 30 - offset
// 	}
// 	for i := offset + 1; i <= offset+limit; i++ {
// 		reportArr = append(reportArr, reportMap[i])
// 	}
// 	return reportArr, nil
// }

func NewMysqlTransactionRepository(db *sql.DB) transaction.TransactionRepository {
	return &mysqlTransactionRepo{
		DB: db,
	}
}
