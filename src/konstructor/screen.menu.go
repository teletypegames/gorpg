package konstructor

import "image/color"

type MenuMapKey string

// Map of menus
type MenuMap map[MenuMapKey]Menu

// Layout of menu
type MenuLayout struct {
	MenuItemFont FontLayout
}

// Menu item
type MenuItem struct {
	ID      string
	Label   string
	Handler func()
}

// Game menu
type Menu struct {
	CurrentSelected int
	Layout          MenuLayout
	MenuItems       []MenuItem
}

// Return with the color of menu item
func (menu *Menu) GetMenuItemColor(i int) color.Color {
	if menu.CurrentSelected == i {
		return menu.Layout.MenuItemFont.SelectedColor
	} else {
		return menu.Layout.MenuItemFont.Color
	}
}
