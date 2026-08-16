package easy_ebitengine

import (
	"git.teletypegames.org/engines/gorpg/src/konstructor"

	"github.com/hajimehoshi/ebiten"
)

func (e *Engine) InventoryUpdate() {
}

func (e *Engine) InventoryDraw(screen *ebiten.Image) {
}

func (e *Engine) AddToInventory(item *konstructor.Item) {
	e.Domain.AddToInventory(item)
}

func (e *Engine) RemoveFromInventory(item *konstructor.Item) {
	e.Domain.RemoveFromInventory(item)
}

func (e *Engine) UseInventoryItem(item *konstructor.Item) {
	e.Domain.UseInventoryItem(item)
}
