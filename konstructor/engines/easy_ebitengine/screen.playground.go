package easy_ebitengine

import (
	"game/konstructor"

	"github.com/hajimehoshi/ebiten"
)

func (e *Engine) PlaygroundUpdate() {
	level := e.Domain.GetLevel(e.KContext.CurrentLevel)
	e.Domain.Process(konstructor.DomainProcessArgs{
		Level:    &level,
		KContext: e.KContext,
	})
}

func (e *Engine) PlaygroundDraw(screen *ebiten.Image) {
	e.PlaygroundBackgroundDraw(screen)
	e.PlaygroundHeaderDraw(screen)
	e.PlaygroundPlatformsDraw(screen)
	e.PlaygroundItemsDraw(screen)
	e.PlaygroundNPCsDraw(screen)
	e.PlaygroundDefaultPlayerDraw(screen)
}

func (e *Engine) PlaygroundBackgroundDraw(screen *ebiten.Image) {
	playground := e.GetPlayground()
	e.PlaygroundAssetDraw(screen, playground.Render, konstructor.Position{})
}

func (e *Engine) GetPlayground() konstructor.Playground {
	level := e.Domain.GetLevel(e.KContext.CurrentLevel)
	return level.Playgrounds[e.KContext.CurrentPlayground]
}

func (e *Engine) PlaygroundAssetDraw(screen *ebiten.Image, render konstructor.Render, position konstructor.Position) {
	geoM := ebiten.GeoM{}
	geoM.Translate(float64(position.X), float64(position.Y))
	geoM.Scale(float64(e.Config.Screen.Scale), float64(e.Config.Screen.Scale))
	screen.DrawImage(e.GetImage(render), &ebiten.DrawImageOptions{
		GeoM: geoM,
	})
}
