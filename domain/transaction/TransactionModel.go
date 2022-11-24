package transaction

type Transactions struct {
	ID              int32  `json:"id"`
	CurrentQuantity int32  `json:"current_quantity"`
	Price           int32  `json:"price"`
	ProductId       string `json:"product_id"`
	Quantity        int32  `json:"quantity"`
	ShopeId         string `json:"shope_id"`
	Type            string `json:"type"`
	// datetime if needed
	SourceFrom string `json:"source_from"`
	SourceTo   string `json:"source_to"`
}

var currId = int32(2)

func getNextTransactionId() int32 {
	currId++
	return int32(currId)
}
func TransactionFactory(currentQuantity int32, price int32, productId string, quantity int32, shopeId string, actionType string, sourceFrom string, sourceTo string) Transactions {
	return Transactions{getNextTransactionId(), currentQuantity, price, productId, quantity, shopeId, actionType, sourceFrom, sourceTo}
}
