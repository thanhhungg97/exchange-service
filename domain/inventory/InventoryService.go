package inventory

type ProductInventory struct {
	RemainQuantity int32
}

func (inventory ProductInventory) IsStocked(quantity int32) bool {
	return inventory.RemainQuantity >= quantity
}

type IInventoryService interface {
	GetAllInventory() *map[string]ProductInventory
	GetInventory(shopeCode string, productCode string) ProductInventory
	DecreaseInventory(shopeCode string, productCode string, quantity int32)
	IncreaseInventory(shopeCode string, productCode string, quantity int32)
}

type ImpInventoryService struct {
}

func (s *ImpInventoryService) GetInventory(shopeCode string, productCode string) ProductInventory {
	s2 := computeKey(shopeCode, productCode)
	value, exists := inventory[s2]
	if exists {
		return value
	}
	return ProductInventory{0}
}
func computeKey(shopeCode string, productCode string) string {
	return shopeCode + "_" + productCode
}
func (s *ImpInventoryService) DecreaseInventory(shopeCode string, productCode string, quantity int32) {
	s2 := computeKey(shopeCode, productCode)
	value, exists := inventory[s2]
	if exists {
		tmp := value.RemainQuantity - quantity
		inventory[s2] = ProductInventory{tmp}
	}
}
func (s *ImpInventoryService) GetAllInventory() *map[string]ProductInventory {
	return &inventory
}
func (s *ImpInventoryService) IncreaseInventory(shopeCode string, productCode string, quantity int32) {
	s2 := computeKey(shopeCode, productCode)
	value, exists := inventory[s2]
	if exists {
		tmp := value.RemainQuantity + quantity
		inventory[s2] = ProductInventory{tmp}
		return
	}
	inventory[s2] = ProductInventory{quantity}

}
