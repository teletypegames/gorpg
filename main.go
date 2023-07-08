package main

import (
	"game/domain"
	"game/konstructor"
	"game/konstructor/engines/easy_ebitengine"
	"image/color"

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
			Multiplayer:       false,
			LiveCount:         3,
		},
		Config: &konstructor.Config{
			Name: "Teletype   Adventure",
			Screen: &konstructor.ScreenConfig{
				Width:  640,
				Height: 480,
				Scale:  2,
			},
			Header: &konstructor.HeaderConfig{
				BackgroundColor: color.RGBA{
					R: 0,
					G: 255,
					B: 0,
					A: 200,
				},
				Height:     20,
				FontLayout: domain.GetHeaderFontLayout(),
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
