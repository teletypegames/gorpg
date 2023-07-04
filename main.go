package main

import (
	"game/domain"
	"game/konstructor"
	"game/konstructor/engine"
	"game/konstructor/entity"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	k := konstructor.Konstructor{
		KContext: &entity.KContext{
			Screen: entity.KContextScreen{
				Type:  "playground",
				Value: "",
			},
			CurrentLevel: 0,
		},
		Domain: &domain.Domain{
			Context: domain.Context{
				PlayerState: domain.PlayerState{
					Name: "Player One",
				},
				User: domain.User{
					Name: "John Doe",
				},
			},
		},
		Controller: &engine.Controller{
			KeyMap: entity.KeyMap{
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
		Settings: &entity.Settings{
			Name: "Game",
			Screen: &entity.ScreenSettings{
				Width:  640,
				Height: 480,
			},
		},
	}
	k.Init()
}
