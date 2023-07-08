package domain

import "game/src/konstructor"

func (d *Domain) AddToInventory(item *konstructor.Item) bool {
	return true
}

func (d *Domain) RemoveFromInventory(item *konstructor.Item) bool {
	return true
}

func (d *Domain) UseInventoryItem(item *konstructor.Item) bool {
	return true
}
