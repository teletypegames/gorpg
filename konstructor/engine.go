package konstructor

import (
	"github.com/hajimehoshi/ebiten"
)

type Engine struct {
	Konstructor *Konstructor
	Logic       *Logic
}

func (g *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Logic.Settings.Screen.Width, g.Logic.Settings.Screen.Height
}

func (g *Engine) Update(screen *ebiten.Image) error {
	g.Logic.Controller.Watch()
	g.Logic.Update(screen)
	return nil
}

func (g *Engine) Draw(screen *ebiten.Image) {
	g.Logic.Draw(screen)
	g.Logic.Controller.Clear()
}
