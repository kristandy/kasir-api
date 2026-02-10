package model

import "time"

type Produk struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

type Transaction struct {
	ID          int                 `json:"id"`
	TotalAmount int                 `json:"total_amount"`
	CreatedAt   time.Time           `json:"created_at"`
	Details     []TransactionDetail `json:"details"`
}

type TransactionDetail struct {
	ID            int    `json:"id"`
	TransactionID int    `json:"transaction_id"`
	ProductID     int    `json:"product_id"`
	ProductName   string `json:"product_name,omitempty"`
	Quantity      int    `json:"quantity"`
	Subtotal      int    `json:"subtotal"`
}

type CheckoutItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CheckoutRequest struct {
	Items []CheckoutItem `json:"items"`
}

type ReportDetail struct {
	TotalRevenue      int               `json:"total_revenue"`
	TotalTransactions int               `json:"total_transactions"`
	ProdukTerlaris    []ProductTerlaris `json:"produk_terlaris"`
}

type ReportRequest struct {
	DateFrom string `json:"date_from"`
	DateTo   string `json:"date_to"`
}

type ProductTerlaris struct {
	ProductName string `json:"product_name"`
	Quantity    int    `json:"qty_sold"`
}
