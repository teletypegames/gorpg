package engine

import (
	"game/konstructor/entity"
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten"
)

func (e *Engine) GetImage(render_options entity.Render) *ebiten.Image {
	img := render_options.GetImage()
	out, err := ebiten.NewImageFromImage(img, 0)
	if err != nil {
		log.Fatal(err)
	}
	return out
}
