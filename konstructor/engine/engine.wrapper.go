package engine

import (
	"game/konstructor/entity"

	"github.com/hajimehoshi/ebiten"
)

type EngineWrapper struct {
	Engine Engine
}

func (ew *EngineWrapper) Init(options entity.EngineOptions) {
	ew.Engine = Engine{
		KContext: options.KContext,
		Domain:   options.Domain,
		Settings: options.Settings,
	}
	ebiten.SetWindowSize(options.Settings.Screen.Width, options.Settings.Screen.Height)
	ebiten.SetWindowTitle(options.Settings.Name)
}

func (ew *EngineWrapper) Run() {
	ebiten.RunGame(&ew.Engine)
}
