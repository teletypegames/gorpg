package easy_ebitengine

import (
	"fmt"
	"game/konstructor"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

func (e *Engine) PlaygroundHeaderDraw(screen *ebiten.Image) {
	e.PlaygroundHeaderBackgroundDraw(screen)
	e.PlaygroundHeaderTextDraw(screen)
}

func (e *Engine) PlaygroundHeaderBackgroundDraw(screen *ebiten.Image) {
	background, _ := ebiten.NewImage(e.Config.Screen.Width, e.Config.Header.Height, ebiten.FilterDefault)
	background.Fill(e.Config.Header.BackgroundColor)
	geoM := ebiten.GeoM{}
	geoM.Scale(float64(e.Config.Screen.Scale), float64(e.Config.Screen.Scale))
	screen.DrawImage(background, &ebiten.DrawImageOptions{GeoM: geoM})
}

func (e *Engine) PlaygroundHeaderTextDraw(screen *ebiten.Image) {
	face := e.Config.Header.FontLayout.GetFontFace(e.Config.Screen.Scale)
	header_text := fmt.Sprint(e.Config.Name, "    LIVES   ", e.KContext.LiveCount, "    LEVEL    ", e.KContext.CurrentLevel)

	offset := getTopOffset(e.Config)
	text.Draw(screen, header_text, face, 10, offset, color.Black)
}

func getTopOffset(config *konstructor.Config) int {
	size := int(config.Header.FontLayout.Size) * config.Screen.Scale
	height := config.Header.Height * config.Screen.Scale
	return (height-size)/2 + size
}
