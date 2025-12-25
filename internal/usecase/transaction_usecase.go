package usecase

import (
	"fmt"

	"github.com/MarcelloBB/plata/internal/dto"
	"github.com/MarcelloBB/plata/internal/model"
	"github.com/MarcelloBB/plata/internal/repository"
	"github.com/MarcelloBB/plata/internal/utils"
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

func (p *TransactionUseCase) CreateTransaction(transaction dto.CreateTransactionRequest) (dto.TransactionResponse, error) {
	transactionDate, transactionDateStr, parseDateErr := utils.NormalizeDate(transaction.Date)
	if parseDateErr != nil {
		return dto.TransactionResponse{}, fmt.Errorf("invalid date: %w", parseDateErr)
	}

	newTransaction := model.Transaction{
		UserID:      transaction.UserID,
		Amount:      transaction.Amount,
		Description: transaction.Description,
		Date:        transactionDate,
		Type:        model.TransactionType(transaction.Type),
	}

	createdTransaction, err := p.repository.CreateTransaction(newTransaction)
	if err != nil {
		return dto.TransactionResponse{}, err
	}

	response := dto.TransactionResponse{
		ID:          createdTransaction.ID,
		UserID:      createdTransaction.UserID,
		Amount:      createdTransaction.Amount,
		Description: createdTransaction.Description,
		Date:        transactionDateStr,
		Type:        string(createdTransaction.Type),
	}

	return response, nil
}
