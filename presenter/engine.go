package presenter

import "github.com/hajimehoshi/ebiten"

const (
	screenWidth  = 320
	screenHeight = 240
)

type GameEngine struct {
	keys []ebiten.Key
}

func (g *GameEngine) Update(*ebiten.Image) error {
	return nil
}

func (g *GameEngine) Draw(screen *ebiten.Image) {

}

func (g *GameEngine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
