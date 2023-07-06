package easy_ebitengine

import (
	"game/konstructor"
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten"
)

func (e *Engine) GetImage(render_options konstructor.Render) *ebiten.Image {
	img := render_options.GetImage()
	out, err := ebiten.NewImageFromImage(img, 0)
	if err != nil {
		log.Fatal(err)
	}
	return out
}
