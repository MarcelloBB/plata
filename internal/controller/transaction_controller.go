package controller

import (
	"net/http"

	"github.com/MarcelloBB/plata/internal/usecase"
	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	transactionUseCase usecase.TransactionUseCase
}

func NewTransactionController(usecase usecase.TransactionUseCase) TransactionController {
	return TransactionController{
		transactionUseCase: usecase,
	}
}

// GetTransactions godoc
// @Summary      List transactions
// @Description  Returns a list of transactions
// @Tags         transaction
// @Produce      json
// @Success      200  {array}   model.Transaction
// @Failure      500  {object}  map[string]string
// @Router       /transaction [get]
func (uc *TransactionController) GetTransactions(c *gin.Context) {
	transactions, err := uc.transactionUseCase.GetTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}
	c.JSON(http.StatusOK, transactions)
}
