package konstructor

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type Logic struct {
	Konstructor *Konstructor
}

func (l *Logic) Update(screen *ebiten.Image) {
}

func (l *Logic) Draw(screen *ebiten.Image) {
	if l.Konstructor.Controller.UpPressed() {
		screen.Fill(color.White)
	}
}
