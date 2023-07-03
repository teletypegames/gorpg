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

func (e *Engine) Update(screen *ebiten.Image) error {
	e.Controller.Watch()
	if e.KContext.Screen.Type == "menu" {
		e.MenuUpdate()
	}
	return nil
}

func (e *Engine) Draw(screen *ebiten.Image) {
	if e.KContext.Screen.Type == "menu" {
		e.MenuDraw(screen)
	}
	if e.Controller.Action3Pressed() {
		os.Exit(1)
	}
	e.Controller.Clear()
}
