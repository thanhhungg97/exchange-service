package inventory

var inventory = make(map[string]ProductInventory)

func LoadInventory() {
	inventory["shope_1_product_1"] = ProductInventory{10}
	inventory["shope_1_product_2"] = ProductInventory{10}
}
