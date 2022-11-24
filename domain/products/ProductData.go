package products

var products = []Products{
	{ID: "1", ProductCode: "product_1", Type: "Bánh", Name: "Bánh quy"},
	{ID: "2", ProductCode: "product_2", Type: "Kẹo", Name: "Kẹo Ngọt"},
}

var shopeProducts = []ShopeProduct{
	{ProductId: "1", ShopeId: "1", Price: 5000},
	{ProductId: "2", ShopeId: "1", Price: 1000},
}
var shopes = []Shopes{
	{ID: "1", ShopeCode: "shope_1"},
	{ID: "2", ShopeCode: "shope_2"},
}
