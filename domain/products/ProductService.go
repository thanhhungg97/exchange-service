package products

type IProductService interface {
	GetProductByProductCode(code string) Products
	GetProductByShopeAndProductCode(shopeCode string, productCode string) ShopeProduct
}
type ImpProductService struct {
}

func (ImpProductService) GetProductByProductCode(productCode string) Products {
	for _, product := range products {
		if product.ProductCode == productCode {
			return product
		}
	}
	return products[0]
}
func (ImpProductService) GetProductByShopeAndProductCode(shopeCode string, productCode string) ShopeProduct {
	var productId = ""
	var shopeId = ""
	for _, product := range products {
		if product.ProductCode == productCode {
			productId = product.ID
		}
	}
	for _, shope := range shopes {
		if shope.ShopeCode == shopeCode {
			shopeId = shope.ID
		}
	}
	for _, tmp := range shopeProducts {
		if tmp.ProductId == productId && tmp.ShopeId == shopeId {
			return tmp
		}
	}

	// throw something
	return shopeProducts[0]
}
