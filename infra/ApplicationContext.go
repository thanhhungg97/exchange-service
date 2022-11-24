package infra

import (
	"awesomeProject/domain/exchange"
	"awesomeProject/domain/inventory"
	"awesomeProject/domain/products"
	"awesomeProject/domain/transaction"
)

type ApplicationContext struct {
	InventoryService   inventory.IInventoryService
	TransactionService transaction.ITransactionService
	ProductService     products.IProductService
	ExchangeService    exchange.IExchangeService
}

func ApplicationContextFactory() ApplicationContext {
	inventoryService := &inventory.ImpInventoryService{}
	productService := &products.ImpProductService{}
	transactionService := &transaction.ImpTransactionService{}
	exchangeService := &exchange.ImpExchangeService{
		InventoryService:   inventoryService,
		ProductService:     productService,
		TransactionService: transactionService}
	return ApplicationContext{inventoryService, transactionService, productService, exchangeService}
}
