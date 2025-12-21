package usecase

import (
	"github.com/MarcelloBB/plata/internal/model"
	"github.com/MarcelloBB/plata/internal/repository"
)

type TransactionUseCase struct {
	repository repository.TransactionRepository
}

func NewTransactionUseCase(repo repository.TransactionRepository) TransactionUseCase {
	return TransactionUseCase{
		repository: repo,
	}
}

func (p *TransactionUseCase) GetTransactions() ([]model.Transaction, error) {
	return p.repository.GetTransactions()
}
