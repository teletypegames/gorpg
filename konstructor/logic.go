package konstructor

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type Logic struct {
	Domain     DomainInterface
	Controller *Controller
	Settings   *Settings
}

func (l *Logic) Update(screen *ebiten.Image) {
}

func (l *Logic) Draw(screen *ebiten.Image) {
	if l.Controller.UpPressed() {
		screen.Fill(color.White)
	}
}
