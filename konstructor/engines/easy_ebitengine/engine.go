package easy_ebitengine

import (
	"game/konstructor"
	"os"

	"github.com/hajimehoshi/ebiten"
)

type Engine struct {
	Domain     konstructor.DomainInterface
	Config     *konstructor.Config
	KContext   *konstructor.KContext
	PressedKey ebiten.Key
}

func (e *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return e.Config.Screen.Width, e.Config.Screen.Height
}

func (e *Engine) Update(screen *ebiten.Image) error {
	e.WatchKeyPress()
	if e.KContext.ScreenTypeIs(konstructor.MenuScreenType) {
		e.MenuUpdate()
	}
	if e.KContext.ScreenTypeIs(konstructor.DialogScreenType) {
		e.DialogUpdate()
	}
	if e.KContext.ScreenTypeIs(konstructor.PlaygroundScreenType) {
		e.PlaygroundUpdate()
	}
	if e.KContext.ScreenTypeIs(konstructor.InventoryScreenType) {
		e.InventoryUpdate()
	}
	return nil
}

func (e *Engine) Draw(screen *ebiten.Image) {
	if e.KContext.ScreenTypeIs(konstructor.MenuScreenType) {
		e.MenuDraw(screen)
	}
	if e.KContext.ScreenTypeIs(konstructor.DialogScreenType) {
		e.DialogDraw(screen)
	}
	if e.KContext.ScreenTypeIs(konstructor.PlaygroundScreenType) {
		e.PlaygroundDraw(screen)
	}
	if e.KContext.ScreenTypeIs(konstructor.InventoryScreenType) {
		e.InventoryDraw(screen)
	}
	if e.Action3Pressed() {
		os.Exit(1)
	}
	e.ClearKeyPresed()
}
