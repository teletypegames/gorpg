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
	e.PlaygroundObjectsDraw(screen)
}

func (e *Engine) PlaygroundBackgroundDraw(screen *ebiten.Image) {
	playground := e.GetPlayground()
	render := playground.Render
	screen.DrawImage(e.GetImage(render), &ebiten.DrawImageOptions{})
}

func (e *Engine) PlaygroundObjectsDraw(screen *ebiten.Image) {
	playground := e.GetPlayground()
	geoM := ebiten.GeoM{}
	for _, object := range playground.Objects {
		render := object.Type.Render
		geoM.Reset()
		geoM.Translate(float64(object.Position.X), float64(object.Position.Y))
		screen.DrawImage(e.GetImage(render), &ebiten.DrawImageOptions{
			GeoM: geoM,
		})
	}
}

func (e *Engine) GetPlayground() konstructor.Playground {
	level := e.Domain.GetLevel(e.KContext.CurrentLevel)
	return level.Playgrounds[e.KContext.CurrentPlayground]
}
