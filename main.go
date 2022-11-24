package main

import (
	"awesomeProject/domain/inventory"
	"awesomeProject/infra"
	"awesomeProject/input/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	inventory.LoadInventory()
	var applicationContext = infra.ApplicationContextFactory()

	exchangeHandler := controllers.ExchangeHandler{ExchangeService: applicationContext.ExchangeService}
	transactionHandler := controllers.TransactionHandler{TransactionService: applicationContext.TransactionService}
	inventoryHandler := controllers.InventoryHandler{InventoryService: applicationContext.InventoryService}

	router := gin.Default()

	router.POST("/sell-products", exchangeHandler.SellProduct)
	router.POST("/exchange-products", exchangeHandler.ExchangeProduct)
	router.POST("/import-products", exchangeHandler.ImportProduct)
	router.GET("/transactions", transactionHandler.GetListTransaction)
	router.GET("/inventory", inventoryHandler.GetCurrentInventory)

	router.Run("localhost:8080")
}
