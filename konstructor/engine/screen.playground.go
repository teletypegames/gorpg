package engine

import (
	"github.com/hajimehoshi/ebiten"
)

func (e *Engine) PlaygroundDraw(screen *ebiten.Image) {
	e.Domain.GetLevel(e.KContext.CurrentLevel)
}

func (e *Engine) PlaygroundUpdate() {
	e.Domain.GetLevel(e.KContext.CurrentLevel)
}
