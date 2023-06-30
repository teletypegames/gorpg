package main

import (
	"game/domain"
	"game/konstructor"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	domain := domain.Domain{
		Context: domain.Context{
			Player: domain.Player{
				Name: "Player One",
			},
			User: domain.User{
				Name: "John Doe",
			},
		},
	}
	konstructor := konstructor.Konstructor{
		Domain: domain,
		Controller: konstructor.Controller{
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
		GameSettings: konstructor.GameSettings{
			Name: "Game",
			ScreenSettings: konstructor.ScreenSettings{
				Width:  640,
				Height: 480,
			},
		},
	}
	konstructor.Init()
}
