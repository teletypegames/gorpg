package domain

import (
	"fmt"
	"game/konstructor/entity"
	"os"
)

func (d *Domain) InitMenu() {
	d.MenuMap = entity.MenuMap{
		"MainMenu": {
			CurrentSelected: 0,
			Layout: entity.MenuLayout{
				MenuItemFont: GetDefaultFontLayout(),
			},
			MenuItems: []entity.MenuItem{
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
			Layout: entity.MenuLayout{
				MenuItemFont: GetDefaultFontLayout(),
			},
			MenuItems: []entity.MenuItem{
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

func (d *Domain) GetMenuMap() entity.MenuMap {
	return d.MenuMap
}

func (d *Domain) GetMenu(name string) entity.Menu {
	value, _ := d.MenuMap[name]
	return value
}

func (d *Domain) SetMenu(name string, menu entity.Menu) {
	d.MenuMap[name] = menu
}

func (d *Domain) CloseGame() {
	fmt.Println("Close game")
}

func (d *Domain) LoadtGame() {
	fmt.Println("Load game")
}

func (d *Domain) SaveGame() {
	fmt.Println("Save game")
}

func (d *Domain) StartGame() {
	fmt.Println("Start game")
}

func (d *Domain) ExitGame() {
	fmt.Println("Exit game")
	os.Exit(1)
}
