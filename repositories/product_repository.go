package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/model"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (reps *ProductRepository) GetProducts(name string) ([]model.Produk, error) {
	args := []any{}
	query := "select id, name, price, stock from products"
	if name != "" {
		query += " where name ILIKE $1"
		args = append(args, "%"+name+"%")
	}
	rows, err := reps.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]model.Produk, 0)
	for rows.Next() {
		var p model.Produk
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (reps *ProductRepository) CreateProduct(product *model.Produk) error {
	query := "insert into products (name, price, stock) values ($1, $2, $3) returning id"
	err := reps.db.QueryRow(query, product.Name, product.Price, product.Stock).Scan(&product.ID)
	return err
}

func (repo *ProductRepository) GetByID(id int) (*model.Produk, error) {
	query := "SELECT id, name, price, stock FROM products WHERE id = $1"

	var p model.Produk
	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock)
	if err == sql.ErrNoRows {
		return nil, errors.New("produk tidak ditemukan")
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (repo *ProductRepository) Update(product *model.Produk) error {
	query := "UPDATE products SET name = $1, price = $2, stock = $3 WHERE id = $4"
	result, err := repo.db.Exec(query, product.Name, product.Price, product.Stock, product.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return nil
}

func (repo *ProductRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return err
}
