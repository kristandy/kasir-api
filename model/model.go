package model

type Produk struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}
