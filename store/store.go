package store

import (
	"strconv"

	"../interfaces"
)

type Store struct {
	interfaces.Inventory
}

func (s *Store) GetOrder(id int) string {
	return "take your order " + strconv.Itoa(id)
}
