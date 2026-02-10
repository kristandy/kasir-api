package services

import (
	"kasir-api/model"
	"kasir-api/repositories"
)

type ProductService struct {
	reps *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{reps: repo}
}

func (s *ProductService) GetProducts(name string) ([]model.Produk, error) {
	return s.reps.GetProducts(name)
}

func (s *ProductService) CreateProduct(data *model.Produk) error {
	return s.reps.CreateProduct(data)
}

func (s *ProductService) GetByID(id int) (*model.Produk, error) {
	return s.reps.GetByID(id)
}

func (s *ProductService) Update(product *model.Produk) error {
	return s.reps.Update(product)
}

func (s *ProductService) Delete(id int) error {
	return s.reps.Delete(id)
}
