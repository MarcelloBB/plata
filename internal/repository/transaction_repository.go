package repository

import (
	"fmt"

	"github.com/MarcelloBB/plata/internal/model"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	connection *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return TransactionRepository{
		connection: db,
	}
}

func (p *TransactionRepository) GetTransactions() ([]model.Transaction, error) {
	transactions := []model.Transaction{}
	err := p.connection.Find(&transactions).Error
	if err != nil {
		fmt.Println("Error executing query:", err)
		return []model.Transaction{}, err
	}

	return transactions, nil
}
