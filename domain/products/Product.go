package products

type Shopes struct {
	ID        string `json:"id"`
	ShopeCode string `json:"shope_code"`
}

type Products struct {
	ID          string `json:"id"`
	ProductCode string `json:"product_code"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}
type ShopeProduct struct {
	ProductId string `json:"product_id"`
	ShopeId   string `json:"shope_id"`
	Price     int32  `json:"price"`
}
