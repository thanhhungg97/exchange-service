package controllers

import (
	"awesomeProject/domain/inventory"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InventoryHandler struct {
	InventoryService inventory.IInventoryService
}

func (context InventoryHandler) GetCurrentInventory(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, context.InventoryService.GetAllInventory())
}
