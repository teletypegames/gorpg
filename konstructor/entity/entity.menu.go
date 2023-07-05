package entity

import "image/color"

type MenuMap map[string]Menu

type MenuLayout struct {
	MenuItemFont FontLayout
}

type MenuItem struct {
	ID      string
	Label   string
	Handler func()
}

type Menu struct {
	CurrentSelected int
	Layout          MenuLayout
	MenuItems       []MenuItem
}

func (menu *Menu) GetMenuItemColor(i int) color.Color {
	if menu.CurrentSelected == i {
		return menu.Layout.MenuItemFont.SelectedColor
	} else {
		return menu.Layout.MenuItemFont.Color
	}
}
