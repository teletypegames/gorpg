package easy_ebitengine

import (
	"game/konstructor"

	"github.com/hajimehoshi/ebiten"
)

type EngineWrapper struct {
	Engine Engine
}

func (ew *EngineWrapper) Init(options konstructor.EngineArgs) {
	ew.Engine = Engine{
		KContext: options.KContext,
		Domain:   options.Domain,
		Config:   options.Config,
	}
	ebiten.SetWindowSize(options.Config.Screen.Width, options.Config.Screen.Height)
	ebiten.SetWindowTitle(options.Config.Name)
}

func (ew *EngineWrapper) Run() {
	ebiten.RunGame(&ew.Engine)
}
