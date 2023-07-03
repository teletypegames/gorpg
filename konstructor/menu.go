package konstructor

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type MenuMap map[string]Menu

type MenuLayout struct {
	Background string
	DPI        float64
	Size       float64
}

type MenuItem struct {
	ID      string
	Label   string
	Handler func()
}

type Menu struct {
	CurrentSelected int
	MenuLayout      MenuLayout
	MenuItems       []MenuItem
}

func (e *Engine) MenuDraw(screen *ebiten.Image) {
	menu := e.MenuMap[e.KContext.Screen.Value]
	face := e.menuGetFontFace(menu)

	for i, menu_item := range menu.MenuItems {
		var c color.Color
		if menu.CurrentSelected == i {
			c = color.RGBA{R: 0, G: 255, B: 0, A: 100}
		} else {
			c = color.White
		}
		text.Draw(screen, menu_item.Label+"\n", face, 8, 24*(i+1), c)
	}
}

func (e *Engine) MenuUpdate() {
	menu := e.MenuMap[e.KContext.Screen.Value]
	if e.Controller.DownPressed() && menu.CurrentSelected != len(menu.MenuItems)-1 {
		menu.CurrentSelected++
	}

	if e.Controller.UpPressed() && menu.CurrentSelected != 0 {
		menu.CurrentSelected--
	}

	e.MenuMap[e.KContext.Screen.Value] = menu
}

func (e *Engine) menuGetFontFace(menu Menu) font.Face {
	tt, _ := opentype.Parse(fonts.MPlus1pRegular_ttf)
	face, _ := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    menu.MenuLayout.DPI,
		DPI:     menu.MenuLayout.Size,
		Hinting: font.HintingVertical,
	})
	return face
}
