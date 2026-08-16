package domain

import "git.teletypegames.org/engines/gorpg/src/konstructor"

func (d *Domain) AddToInventory(item *konstructor.Item) bool {
	return true
}

func (d *Domain) RemoveFromInventory(item *konstructor.Item) bool {
	return true
}

func (d *Domain) UseInventoryItem(item *konstructor.Item) bool {
	return true
}
