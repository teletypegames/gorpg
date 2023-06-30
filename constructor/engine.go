package constructor

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type GameEngine struct {
	Constructor Constructor
}

func (g *GameEngine) Update(screen *ebiten.Image) error {
	g.Constructor.Controller.Watch()
	return nil
}

func (g *GameEngine) Draw(screen *ebiten.Image) {
	if g.Constructor.Controller.UpPressed() {
		screen.Fill(color.White)
	}
}

func (g *GameEngine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Constructor.GameSettings.ScreenSettings.Width, g.Constructor.GameSettings.ScreenSettings.Height
}
