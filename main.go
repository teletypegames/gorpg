package main

import (
	"game/constructor"
	"game/domain"

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
	constructor := constructor.Constructor{
		Domain: domain,
		Keyboard: constructor.Keyboard{
			KeyMap: constructor.KeyMap{
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
	constructor.Init()
}
