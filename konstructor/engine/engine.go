package engine

import (
	"game/konstructor/entity"
	"os"

	"github.com/hajimehoshi/ebiten"
)

type Engine struct {
	Domain     entity.DomainInterface
	Settings   *entity.Settings
	KContext   *entity.KContext
	PressedKey ebiten.Key
}

func (e *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return e.Settings.Screen.Width, e.Settings.Screen.Height
}

func (e *Engine) Update(screen *ebiten.Image) error {
	e.WatchKeyPress()
	if e.KContext.ScreenTypeIs(entity.MenuScreenType) {
		e.MenuUpdate()
	}
	if e.KContext.ScreenTypeIs(entity.DialogScreenType) {
		e.DialogUpdate()
	}
	if e.KContext.ScreenTypeIs(entity.PlaygroundScreenType) {
		e.PlaygroundUpdate()
	}
	if e.KContext.ScreenTypeIs(entity.InventoryScreenType) {
		e.InventoryUpdate()
	}
	return nil
}

func (e *Engine) Draw(screen *ebiten.Image) {
	if e.KContext.ScreenTypeIs(entity.MenuScreenType) {
		e.MenuDraw(screen)
	}
	if e.KContext.ScreenTypeIs(entity.DialogScreenType) {
		e.DialogDraw(screen)
	}
	if e.KContext.ScreenTypeIs(entity.PlaygroundScreenType) {
		e.PlaygroundDraw(screen)
	}
	if e.KContext.ScreenTypeIs(entity.InventoryScreenType) {
		e.InventoryDraw(screen)
	}
	if e.Action3Pressed() {
		os.Exit(1)
	}
	e.ClearKeyPresed()
}
