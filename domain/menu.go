package domain

import (
	"game/konstructor"
	"os"
)

func (d Domain) GetMenuMap() konstructor.MenuMap {

	return map[string]konstructor.Menu{
		"MainMenu": {
			MenuItems: []konstructor.MenuItem{
				{
					ID:      "new",
					Label:   "New Game",
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
	}
}

func (d *Domain) CloseGame() {
}

func (d *Domain) LoadtGame() {
}

func (d *Domain) SaveGame() {
}

func (d *Domain) StartGame() {
}

func (d *Domain) ExitGame() {
	os.Exit(1)
}
