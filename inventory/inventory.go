package inventory

import (
	"strconv"

	"../interfaces"
)

type Inventory struct {
	interfaces.Store
}

func (i *Inventory) StockLevels(productID int) string {
	return "we have 300 " + strconv.Itoa(productID)
}
