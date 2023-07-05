package engine

import "game/konstructor/entity"

func (e *Engine) AddToInventory(item *entity.Item) {
	e.Domain.AddToInventory(item)
}

func (e *Engine) RemoveFromInventory(item *entity.Item) {
	e.Domain.RemoveFromInventory(item)
}

func (e *Engine) UseInventoryItem(item *entity.Item) {
	e.Domain.UseInventoryItem(item)
}
