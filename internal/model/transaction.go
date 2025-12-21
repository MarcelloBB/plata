package model

import "time"

type TransactionType string

const (
	Expense TransactionType = "expense"
	Income  TransactionType = "income"
)

type Transaction struct {
	// example: 1
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`
	// example: 1
	UserID int  `gorm:"not null;index" json:"user_id"`
	User   User `gorm:"foreignKey:UserID" json:"-"`
	// example: 100.50
	Amount float64 `json:"amount"`
	// example: Payment for services
	Description string `gorm:"size:75" json:"description"`
	// example: 2024-01-15T14:30:00Z
	Date time.Time `json:"date"`
	// example: expense
	Type TransactionType `gorm:"type:VARCHAR(10)" json:"type"`
}
