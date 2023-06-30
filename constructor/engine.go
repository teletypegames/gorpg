package constructor

import (
	"game/domain"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

type GameEngine struct {
	Context  domain.Context
	Keyboard Keyboard
}

func (g *GameEngine) Update(screen *ebiten.Image) error {
	g.Keyboard.Watch()
	return nil
}

func (g *GameEngine) Draw(screen *ebiten.Image) {
	if g.Keyboard.UpPressed() {
		screen.Fill(color.White)
	}
}

func (g *GameEngine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
