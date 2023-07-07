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
	e.PlaygroundPlatformsDraw(screen)
	e.PlaygroundItemsDraw(screen)
	e.PlaygroundNPCsDraw(screen)
	e.PlaygroundDefaultPlayerDraw(screen)
}

func (e *Engine) PlaygroundBackgroundDraw(screen *ebiten.Image) {
	playground := e.GetPlayground()
	render := playground.Render
	screen.DrawImage(e.GetImage(render), &ebiten.DrawImageOptions{})
}

func (e *Engine) PlaygroundPlatformsDraw(screen *ebiten.Image) {
	playground := e.GetPlayground()
	geoM := ebiten.GeoM{}
	for _, platform := range playground.Platforms {
		render := platform.Type.Render
		geoM.Reset()
		geoM.Translate(float64(platform.Position.X), float64(platform.Position.Y))
		screen.DrawImage(e.GetImage(render), &ebiten.DrawImageOptions{
			GeoM: geoM,
		})
	}
}

func (e *Engine) PlaygroundItemsDraw(screen *ebiten.Image) {
	playground := e.GetPlayground()
	geoM := ebiten.GeoM{}
	for _, platform := range playground.Items {
		render := platform.Type.Render
		geoM.Reset()
		geoM.Translate(float64(platform.Position.X), float64(platform.Position.Y))
		screen.DrawImage(e.GetImage(render), &ebiten.DrawImageOptions{
			GeoM: geoM,
		})
	}
}

func (e *Engine) PlaygroundNPCsDraw(screen *ebiten.Image) {
	playground := e.GetPlayground()
	geoM := ebiten.GeoM{}
	for _, platform := range playground.NPCs {
		render := platform.Type.Render
		geoM.Reset()
		geoM.Translate(float64(platform.Position.X), float64(platform.Position.Y))
		screen.DrawImage(e.GetImage(render), &ebiten.DrawImageOptions{
			GeoM: geoM,
		})
	}
}

func (e *Engine) PlaygroundDefaultPlayerDraw(screen *ebiten.Image) {
	player := e.Domain.GetDefaultPlayer()
	geoM := ebiten.GeoM{}
	render := player.Type.Render
	geoM.Reset()
	geoM.Translate(float64(player.Position.X), float64(player.Position.Y))
	screen.DrawImage(e.GetImage(render), &ebiten.DrawImageOptions{
		GeoM: geoM,
	})
}

func (e *Engine) GetPlayground() konstructor.Playground {
	level := e.Domain.GetLevel(e.KContext.CurrentLevel)
	return level.Playgrounds[e.KContext.CurrentPlayground]
}
