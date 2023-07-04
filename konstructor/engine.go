package konstructor

import (
	"os"

	"github.com/hajimehoshi/ebiten"
)

type Engine struct {
	Domain     DomainInterface
	Controller *Controller
	Settings   *Settings
	KContext   *KContext
	MenuMap    MenuMap
}

func (e *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return e.Settings.Screen.Width, e.Settings.Screen.Height
}

func (e *Engine) ScreenTypeIs(name string) bool {
	return e.KContext.Screen.Type == name
}

func (e *Engine) Update(screen *ebiten.Image) error {
	e.Controller.Watch()
	if e.ScreenTypeIs("menu") {
		e.MenuUpdate()
	}
	if e.ScreenTypeIs("dialog") {
		e.DialogUpdate()
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
	if e.Controller.Action3Pressed() {
		os.Exit(1)
	}
	e.Controller.Clear()
}
