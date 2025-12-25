package controller

import (
	"net/http"

	"github.com/MarcelloBB/plata/internal/dto"
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

// PostTransactions godoc
// @Summary      Create transaction
// @Description  Create a transaction
// @Tags         transaction
// @Produce      json
// @Success      200  {object}   model.Transaction
// @Failure      500  {object}  map[string]string
// @Router       /transaction [post]
func (uc *TransactionController) PostTransactions(c *gin.Context) {
	var newTransaction dto.CreateTransactionRequest

	if err := c.ShouldBindJSON(&newTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := uc.transactionUseCase.CreateTransaction(newTransaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}
