package controllers

import (
	"awesomeProject/domain/exchange"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SellITemDto struct {
	ShopeCode   string `json:"shope_code"`
	ProductCode string `json:"product_code"`
	Quantity    int32  `json:"quantity"`
}
type SellItemResponse struct {
	IsSuccess bool  `json:"is_success"`
	Amount    int32 `json:"amount"`
}

type ExchangeProductDto struct {
	ShopeCodeFrom string `json:"shope_code_from"`
	ShopeCodeTo   string `json:"shope_code_to"`
	ProductCode   string `json:"product_code"`
	Quantity      int32  `json:"quantity"`
}

type ImportProductDto struct {
	ShopeCode   string `json:"shope_code"`
	ProductCode string `json:"product_code"`
	Quantity    int32  `json:"quantity"`
}

type ExchangeHandler struct {
	ExchangeService exchange.IExchangeService
}

func (s ExchangeHandler) SellProduct(c *gin.Context) {
	var dto SellITemDto

	if err := c.BindJSON(&dto); err != nil {
		return
	}
	data := s.ExchangeService.ExportProduct(exchange.ExportCommand{
		ProductCode: dto.ProductCode,
		ShopeCode:   dto.ShopeCode,
		Quantity:    dto.Quantity,
		ExportTo:    "EndUser",
	})
	c.IndentedJSON(http.StatusCreated, &SellItemResponse{IsSuccess: data.IsSuccess, Amount: data.Amount})

}
func (s ExchangeHandler) ExchangeProduct(c *gin.Context) {
	var dto ExchangeProductDto

	if err := c.BindJSON(&dto); err != nil {
		return
	}
	data := s.ExchangeService.ExportProduct(exchange.ExportCommand{
		ProductCode: dto.ProductCode,
		ShopeCode:   dto.ShopeCodeFrom,
		Quantity:    dto.Quantity,
		ExportTo:    dto.ShopeCodeTo,
	})
	if !data.IsSuccess {
		c.IndentedJSON(http.StatusBadRequest, false)
		return
	}
	s.ExchangeService.ImportProduct(exchange.ImportCommand{
		ProductCode: dto.ProductCode,
		ShopeCode:   dto.ShopeCodeTo,
		Quantity:    dto.Quantity,
		ImportFrom:  dto.ShopeCodeFrom,
	})

	c.IndentedJSON(http.StatusCreated, true)
}

func (s ExchangeHandler) ImportProduct(c *gin.Context) {
	var dto ImportProductDto

	if err := c.BindJSON(&dto); err != nil {
		return
	}

	data := s.ExchangeService.ImportProduct(exchange.ImportCommand{
		ProductCode: dto.ProductCode,
		ShopeCode:   dto.ShopeCode,
		Quantity:    dto.Quantity,
		ImportFrom:  "WareHouse",
	})
	c.IndentedJSON(http.StatusCreated, data)

}
