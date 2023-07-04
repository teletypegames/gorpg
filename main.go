package main

import (
	"game/domain"
	"game/konstructor"

	"github.com/hajimehoshi/ebiten"
)

func main() {

	k := konstructor.Konstructor{
		KContext: &konstructor.KContext{
			Screen: konstructor.KContextScreen{
				Type:  "dialog",
				Value: "TestDialog",
			},
		},
		Domain: &domain.Domain{
			Context: domain.Context{
				Player: domain.Player{
					Name: "Player One",
				},
				User: domain.User{
					Name: "John Doe",
				},
			},
		},
		Controller: &konstructor.Controller{
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
		Settings: &konstructor.Settings{
			Name: "Game",
			Screen: &konstructor.ScreenSettings{
				Width:  640,
				Height: 480,
			},
		},
	}
	k.Init()
}
