package main

import (
	"game/domain"
	"game/konstructor"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

func main() {

	d := domain.Domain{
		Context: domain.Context{
			Player: domain.Player{
				Name: "Player One",
			},
			User: domain.User{
				Name: "John Doe",
			},
		},
	}

	k := konstructor.Konstructor{
		KContext: &konstructor.KContext{
			Screen: konstructor.KContextScreen{
				Type:  "menu",
				Value: "MainMenu",
			},
		},
		Domain: &d,
		MenuMap: konstructor.MenuMap{
			"MainMenu": {
				CurrentSelected: 0,
				MenuLayout: konstructor.MenuLayout{
					DPI:           72,
					Size:          24,
					Color:         color.White,
					SelectedColor: color.RGBA{R: 0, G: 255, B: 0, A: 100},
				},
				MenuItems: []konstructor.MenuItem{
					{
						ID:      "start",
						Label:   "Start Game",
						Handler: d.StartGame,
					},
					{
						ID:      "load",
						Label:   "Load Game",
						Handler: d.LoadtGame,
					},
					{
						ID:      "exit",
						Label:   "Exit Game",
						Handler: d.ExitGame,
					},
				},
			},
			"GameMenu": {
				CurrentSelected: 0,
				MenuLayout: konstructor.MenuLayout{
					DPI:           72,
					Size:          24,
					Color:         color.White,
					SelectedColor: color.RGBA{R: 0, G: 255, B: 0, A: 100},
				},
				MenuItems: []konstructor.MenuItem{
					{
						ID:      "save",
						Label:   "Save Game",
						Handler: d.SaveGame,
					},
					{
						ID:      "close",
						Label:   "Close Game",
						Handler: d.CloseGame,
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
	k.Init()
}
