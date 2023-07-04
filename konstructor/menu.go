package konstructor

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

type MenuMap map[string]Menu

type MenuLayout struct {
	Background   string
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

func (e *Engine) MenuDraw(screen *ebiten.Image) {
	menu := e.Domain.GetMenu(e.KContext.Screen.Value)
	face := GetFontFace(menu.Layout.MenuItemFont)

	for i, menu_item := range menu.MenuItems {
		color := menu.GetMenuItemColor(i)
		offset := int(menu.Layout.MenuItemFont.Size) * (i + 1)
		text.Draw(screen, menu_item.Label+"\n", face, 8, offset, color)
	}
}

func (e *Engine) MenuUpdate() {
	menu := e.Domain.GetMenu(e.KContext.Screen.Value)

	if e.Controller.UpPressed() && menu.CurrentSelected != 0 {
		menu.CurrentSelected--
	}

	if e.Controller.DownPressed() && menu.CurrentSelected != len(menu.MenuItems)-1 {
		menu.CurrentSelected++
	}

	if e.Controller.Action0Pressed() {
		menu.MenuItems[menu.CurrentSelected].Handler()
	}

	e.Domain.SetMenu(e.KContext.Screen.Value, menu)
}
