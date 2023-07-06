package easy_ebitengine

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

func (e *Engine) MenuUpdate() {
	menu := e.Domain.GetMenu(e.KContext.ActiveMenu)

	if e.UpPressed() && menu.CurrentSelected != 0 {
		menu.CurrentSelected--
	}

	if e.DownPressed() && menu.CurrentSelected != len(menu.MenuItems)-1 {
		menu.CurrentSelected++
	}

	if e.Action0Pressed() {
		menu.MenuItems[menu.CurrentSelected].Handler()
	}

	e.Domain.SetMenu(e.KContext.ActiveMenu, menu)
}

func (e *Engine) MenuDraw(screen *ebiten.Image) {
	menu := e.Domain.GetMenu(e.KContext.ActiveMenu)
	face := menu.Layout.MenuItemFont.GetFontFace()

	for i, menu_item := range menu.MenuItems {
		color := menu.GetMenuItemColor(i)
		offset := int(menu.Layout.MenuItemFont.Size) * (i + 1)
		text.Draw(screen, menu_item.Label+"\n", face, 8, offset, color)
	}
}
