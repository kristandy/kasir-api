package repositories

import (
	"database/sql"
	"kasir-api/model"
	"time"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (reps *ReportRepository) FetchReport(dateFrom, dateTo time.Time) (*model.ReportDetail, error) {
	var report model.ReportDetail
	rp, err := reps.db.Begin()
	if err != nil {
		return nil, err
	}
	defer rp.Rollback()

	query := `
	select
		sum(total_amount) as total_revenue,
		count(id) as total_transactions
	from transactions
	where created_at between $1 and $2`

	err = reps.db.QueryRow(query, dateFrom, dateTo).Scan(&report.TotalRevenue, &report.TotalTransactions)
	if err != nil {
		return nil, err
	}

	queryTerlaris := `
	select
		A.Name as product_name
		,sum(B.quantity) as qty_sold
	from products A
	join transaction_details B on A.id = B.product_id
	where B.quantity = (
		select max(quantity)
		from transaction_details
	)
	group by A.Name`

	var product_name string
	var qty_sold int

	err = reps.db.QueryRow(queryTerlaris, dateFrom, dateTo).Scan(&product_name, &qty_sold)
	if err == nil {
		report.ProdukTerlaris = []model.TransactionDetail{
			{
				ProductName: product_name,
				Quantity:    qty_sold,
			},
		}
	} else {
		report.ProdukTerlaris = []model.TransactionDetail{}
	}
	return &report, nil
}
