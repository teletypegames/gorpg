package domain

import "game/konstructor/entity"

func (d *Domain) AddToInventory(item *entity.Item) bool {
	return true
}

func (d *Domain) RemoveFromInventory(item *entity.Item) bool {
	return true
}

func (d *Domain) UseInventoryItem(item *entity.Item) bool {
	return true
}
