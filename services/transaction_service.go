package services

import (
	"kasir-api/model"
	"kasir-api/repositories"
)

type TransactionService struct {
	reps *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{reps: repo}
}

func (s *TransactionService) Checkout(items []model.CheckoutItem) (*model.Transaction, error) {
	return s.reps.CreateTransaction(items)
}
