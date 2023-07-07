package main

import (
	"game/domain"
	"game/konstructor"
	"game/konstructor/engines/easy_ebitengine"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	k := konstructor.Konstructor{
		EngineWrapper: &easy_ebitengine.EngineWrapper{},
		Domain:        &domain.Domain{},
		KContext: &konstructor.KContext{
			ScreenType:        konstructor.PlaygroundScreenType,
			ActiveMenu:        domain.MainMenu,
			ActiveDialog:      domain.DialogTest,
			CurrentLevel:      0,
			CurrentPlayground: 0,
		},
		Config: &konstructor.Config{
			Name: "Game",
			Screen: &konstructor.ScreenConfig{
				Width:  320,
				Height: 240,
				Scale:  2,
			},
			KeyMap: konstructor.KeyMap{
				Up:      ebiten.KeyUp,
				Down:    ebiten.KeyDown,
				Right:   ebiten.KeyRight,
				Left:    ebiten.KeyLeft,
				Action0: ebiten.KeySpace,
				Action1: ebiten.KeyAlt,
				Action2: ebiten.KeyControl,
				Action3: ebiten.KeyEscape,
			},
		},
	}
	k.Init()
}
