package controllers

import (
	"awesomeProject/domain/transaction"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionHandler struct {
	TransactionService transaction.ITransactionService
}

func (context TransactionHandler) GetListTransaction(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, context.TransactionService.ListTransaction())
}
