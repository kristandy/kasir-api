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

	query := `
	select
		coalesce(sum(total_amount),0) as total_revenue,
		count(id) as total_transactions
	from transactions
	where created_at between $1 and $2`

	err := reps.db.QueryRow(query, dateFrom, dateTo).Scan(&report.TotalRevenue, &report.TotalTransactions)
	if err != nil {
		return nil, err
	}

	queryTerlaris := `
	select
		A.Name as product_name
		,sum(B.quantity) as qty_sold
	from products A
	join transaction_details B on A.id = B.product_id
	join transactions C on B.transaction_id = C.id
	where C.created_at between $1 and $2
	group by A.Name
	order by qty_sold desc
	limit 1`

	var product_name string
	var qty_sold int

	err = reps.db.QueryRow(queryTerlaris, dateFrom, dateTo).Scan(&product_name, &qty_sold)
	if err == nil {
		report.ProdukTerlaris = []model.ProductTerlaris{
			{
				ProductName: product_name,
				Quantity:    qty_sold,
			},
		}
	} else {
		report.ProdukTerlaris = []model.ProductTerlaris{}
	}
	return &report, nil
}
