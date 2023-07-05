package engine

import (
	"game/konstructor/entity"

	"github.com/hajimehoshi/ebiten"
)

func (e *Engine) PlaygroundDraw(screen *ebiten.Image) {
	e.Domain.GetLevel(e.KContext.CurrentLevel)
}

func (e *Engine) PlaygroundUpdate() {
	level := e.Domain.GetLevel(e.KContext.CurrentLevel)
	e.Domain.Process(entity.DomainProcessOptions{
		Level:    &level,
		KContext: e.KContext,
	})
}
