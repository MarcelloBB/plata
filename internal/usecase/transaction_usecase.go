package usecase

import (
	"fmt"
	"time"

	"github.com/MarcelloBB/plata/internal/dto"
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

func (p *TransactionUseCase) CreateTransaction(transaction dto.CreateTransactionRequest) (dto.TransactionResponse, error) {
	transactionDate := transaction.Date
	if transactionDate == "" {
		transactionDate = time.Now().Format(time.RFC3339)
	}
	parsedDate, err := time.Parse(time.RFC3339, transactionDate)
	if err != nil {
		return dto.TransactionResponse{}, fmt.Errorf("invalid date format: %w", err)
	}

	newTransaction := model.Transaction{
		UserID:      transaction.UserID,
		Amount:      transaction.Amount,
		Description: transaction.Description,
		Date:        parsedDate,
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
		Date:        createdTransaction.Date.Format(time.RFC3339),
		Type:        string(createdTransaction.Type),
	}

	return response, nil
}
