package konstructor

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type GameEngine struct {
	Konstructor Konstructor
}

func (g *GameEngine) Update(screen *ebiten.Image) error {
	g.Konstructor.Controller.Watch()
	return nil
}

func (g *GameEngine) Draw(screen *ebiten.Image) {
	if g.Konstructor.Controller.UpPressed() {
		screen.Fill(color.White)
	}
}

func (g *GameEngine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Konstructor.GameSettings.ScreenSettings.Width, g.Konstructor.GameSettings.ScreenSettings.Height
}
