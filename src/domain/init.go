package domain

import (
	"game/src/konstructor"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

func (d *Domain) GetInitialKContext() *konstructor.KContext {
	return &konstructor.KContext{
		ScreenType:        konstructor.PlaygroundScreenType,
		ActiveMenu:        MainMenu,
		ActiveDialog:      DialogTest,
		CurrentLevel:      0,
		CurrentPlayground: 0,
		Multiplayer:       false,
		LiveCount:         3,
	}
}

func (d *Domain) GetInitialConfig() *konstructor.Config {
	return &konstructor.Config{
		Name: "Teletype   Adventure",
		Screen: &konstructor.ScreenConfig{
			Width:  konstructor.VGAResolutionSet.Width,
			Height: konstructor.VGAResolutionSet.Height,
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
			FontLayout: GetHeaderFontLayout(),
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
	}
}
