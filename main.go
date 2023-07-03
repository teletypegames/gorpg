package main

import (
	"game/domain"
	"game/konstructor"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	konstructor := konstructor.Konstructor{
		KContext: &konstructor.KContext{
			Screen: konstructor.KContextScreen{
				Type:  "menu",
				Value: "MainMenu",
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
		MenuMap: konstructor.MenuMap{
			"MainMenu": {
				CurrentSelected: 0,
				MenuLayout: konstructor.MenuLayout{
					DPI:  72,
					Size: 24,
				},
				MenuItems: []konstructor.MenuItem{
					{
						ID:    "new",
						Label: "New Game",
					},
					{
						ID:    "load",
						Label: "Load Game",
					},
					{
						ID:    "exit",
						Label: "Exit Game",
					},
				},
			},
			"GameMenu": {
				CurrentSelected: 0,
				MenuLayout: konstructor.MenuLayout{
					DPI:  72,
					Size: 24,
				},
				MenuItems: []konstructor.MenuItem{
					{
						ID:    "save",
						Label: "Save Game",
					},
					{
						ID:    "close",
						Label: "Close Game",
					},
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
	konstructor.Init()
}
