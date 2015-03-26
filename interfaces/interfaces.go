package interfaces

type Store interface {
	GetOrder(id int) string
}

type Inventory interface {
	StockLevels(productID int) string
}
