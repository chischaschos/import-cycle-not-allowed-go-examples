package main

import (
	"./inventory"
	"./store"
	"fmt"
)

func main() {
	store := &store.Store{}
	inventory := &inventory.Inventory{store}
	store.Inventory = inventory

	fmt.Println(store.GetOrder(2))
	fmt.Println(store.StockLevels(20))

	fmt.Println(inventory.GetOrder(2))
	fmt.Println(inventory.StockLevels(20))
}
