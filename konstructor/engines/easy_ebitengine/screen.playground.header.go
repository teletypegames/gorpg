package easy_ebitengine

import (
	"game/konstructor"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

func (e *Engine) PlaygroundHeaderDraw(screen *ebiten.Image) {
	header_bg, _ := ebiten.NewImage(e.Config.Screen.Width, 20, ebiten.FilterDefault)
	header_bg.Fill(color.RGBA{
		R: 0,
		G: 255,
		B: 0,
		A: 200,
	})

	fl := konstructor.FontLayout{
		Path: "assets/fonts/ArcadeClassic.ttf",
		DPI:  72,
		Size: 24,
	}

	face := fl.GetFontFace()

	geoM := ebiten.GeoM{}
	geoM.Scale(float64(e.Config.Screen.Scale), float64(e.Config.Screen.Scale))
	screen.DrawImage(header_bg, &ebiten.DrawImageOptions{
		GeoM: geoM,
	})
	text.Draw(screen, e.Config.Name, face, 10, 20, color.Black)

}
