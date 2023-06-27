package domain

type MenuItem struct {
	ID      string
	Label   string
	Handler func()
}

type Menu struct {
	MenuItems []MenuItem
}

type MenuManager struct {
	MainMenu Menu
	GameMenu Menu
}

func createMenu(menuItems []MenuItem) Menu {
	menu := Menu{
		MenuItems: menuItems,
	}
	return menu
}

func (mm MenuManager) Init() {

	mo := MenuOperation{}

	mm.MainMenu = createMenu([]MenuItem{
		{
			ID:      "new",
			Label:   "New Game",
			Handler: mo.StartGame,
		},
		{
			ID:      "load",
			Label:   "Load Game",
			Handler: mo.LoadtGame,
		},
	})

	mm.GameMenu = createMenu([]MenuItem{
		{
			ID:      "save",
			Label:   "Save Game",
			Handler: mo.SaveGame,
		},
		{
			ID:      "close",
			Label:   "Close Game",
			Handler: mo.CloseGame,
		},
	})
}
