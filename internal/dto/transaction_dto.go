package dto

type CreateTransactionRequest struct {
	UserID      int     `json:"user_id" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"` // greater than 0
	Description string  `json:"description" binding:"required,max=75"`
	Date        string  `json:"date"`
	Type        string  `json:"type" binding:"required,oneof=expense income"`
}

type TransactionResponse struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
	Type        string  `json:"type"`
}
