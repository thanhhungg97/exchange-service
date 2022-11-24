package exchange

import (
	"awesomeProject/domain/inventory"
	"awesomeProject/domain/products"
	"awesomeProject/domain/transaction"
)

type IExchangeService interface {
	ImportProduct(req ImportCommand) ImportResponse
	ExportProduct(req ExportCommand) ExportResponse
}

type ExportResponse struct {
	IsSuccess     bool
	TransactionId int32
	Amount        int32
}

type ImportResponse struct {
	IsSuccess     bool
	TransactionId int32
}
type ImpExchangeService struct {
	InventoryService   inventory.IInventoryService
	ProductService     products.IProductService
	TransactionService transaction.ITransactionService
}

// in one transaction
// currently support 1 thread.

func (r *ImpExchangeService) ExportProduct(exportReq ExportCommand) ExportResponse {
	currentInventory := r.InventoryService.GetInventory(exportReq.ShopeCode, exportReq.ProductCode)
	if !currentInventory.IsStocked(exportReq.Quantity) {
		return ExportResponse{false, 0, 0}
	}
	shopeProduct := r.ProductService.GetProductByShopeAndProductCode(exportReq.ShopeCode, exportReq.ProductCode)
	transactions := transaction.TransactionFactory(currentInventory.RemainQuantity-exportReq.Quantity, shopeProduct.Price, shopeProduct.ShopeId, exportReq.Quantity, shopeProduct.ShopeId, "export", "", exportReq.ExportTo)
	r.InventoryService.DecreaseInventory(exportReq.ShopeCode, exportReq.ProductCode, exportReq.Quantity)
	transaction.GlobalTransTransactions = r.TransactionService.InsertTransaction(transactions)
	return ExportResponse{true, transactions.ID, exportReq.Quantity * shopeProduct.Price}
}

func (r *ImpExchangeService) ImportProduct(importReq ImportCommand) ImportResponse {
	currentInventory := r.InventoryService.GetInventory(importReq.ShopeCode, importReq.ProductCode)
	shopeProduct := r.ProductService.GetProductByShopeAndProductCode(importReq.ShopeCode, importReq.ProductCode)
	transactions := transaction.TransactionFactory(currentInventory.RemainQuantity+importReq.Quantity, shopeProduct.Price, shopeProduct.ShopeId, importReq.Quantity, shopeProduct.ShopeId, "import", importReq.ImportFrom, "")
	transaction.GlobalTransTransactions = r.TransactionService.InsertTransaction(transactions)
	r.InventoryService.IncreaseInventory(importReq.ShopeCode, importReq.ProductCode, importReq.Quantity)
	return ImportResponse{IsSuccess: true, TransactionId: transactions.ID}
}
