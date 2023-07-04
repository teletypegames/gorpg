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
	KeyMap     entity.KeyMap
	PressedKey ebiten.Key
}

func (e *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return e.Settings.Screen.Width, e.Settings.Screen.Height
}

func (e *Engine) ScreenTypeIs(name string) bool {
	return e.KContext.Screen.Type == name
}

func (e *Engine) Update(screen *ebiten.Image) error {
	e.WatchKeyPress()
	if e.ScreenTypeIs("menu") {
		e.MenuUpdate()
	}
	if e.ScreenTypeIs("dialog") {
		e.DialogUpdate()
	}
	if e.ScreenTypeIs("playground") {
		e.PlaygroundUpdate()
	}
	return nil
}

func (e *Engine) Draw(screen *ebiten.Image) {
	if e.ScreenTypeIs("menu") {
		e.MenuDraw(screen)
	}
	if e.ScreenTypeIs("dialog") {
		e.DialogDraw(screen)
	}
	if e.ScreenTypeIs("playground") {
		e.PlaygroundDraw(screen)
	}
	if e.Action3Pressed() {
		os.Exit(1)
	}
	e.ClearKeyPresed()
}
