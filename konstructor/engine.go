package konstructor

import (
	"github.com/hajimehoshi/ebiten"
)

type Engine struct {
	Konstructor *Konstructor
	Logic       *Logic
}

func (g *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Konstructor.Settings.Screen.Width, g.Konstructor.Settings.Screen.Height
}

func (g *Engine) Update(screen *ebiten.Image) error {
	g.Konstructor.Controller.Watch()
	g.Logic.Update(screen)
	return nil
}

func (g *Engine) Draw(screen *ebiten.Image) {
	g.Logic.Draw(screen)
	g.Konstructor.Controller.Clear()
}
