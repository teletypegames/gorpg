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
	if e.KContext.ScreenTypeIs("menu") {
		e.MenuUpdate()
	}
	if e.KContext.ScreenTypeIs("dialog") {
		e.DialogUpdate()
	}
	if e.KContext.ScreenTypeIs("playground") {
		e.PlaygroundUpdate()
	}
	if e.KContext.ScreenTypeIs("inventory") {
		e.InventoryUpdate()
	}
	return nil
}

func (e *Engine) Draw(screen *ebiten.Image) {
	if e.KContext.ScreenTypeIs("menu") {
		e.MenuDraw(screen)
	}
	if e.KContext.ScreenTypeIs("dialog") {
		e.DialogDraw(screen)
	}
	if e.KContext.ScreenTypeIs("playground") {
		e.PlaygroundDraw(screen)
	}
	if e.KContext.ScreenTypeIs("inventory") {
		e.InventoryDraw(screen)
	}
	if e.Action3Pressed() {
		os.Exit(1)
	}
	e.ClearKeyPresed()
}
