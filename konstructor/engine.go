package konstructor

import (
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Engine struct {
	Domain     DomainInterface
	Controller *Controller
	Settings   *Settings
	KContext   *KContext
	MenuMap    MenuMap
}

func (e *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return e.Settings.Screen.Width, e.Settings.Screen.Height
}

func (e *Engine) Update(screen *ebiten.Image) error {
	e.Controller.Watch()
	if e.KContext.Screen.Type == "menu" {
	}
	return nil
}

func (e *Engine) Draw(screen *ebiten.Image) {
	if e.KContext.Screen.Type == "menu" {
		menu := e.MenuMap[e.KContext.Screen.Value]

		if e.Controller.DownPressed() && menu.CurrentSelected != len(menu.MenuItems)-1 {
			menu.CurrentSelected++
		}

		if e.Controller.UpPressed() && menu.CurrentSelected != 0 {
			menu.CurrentSelected--
		}

		e.MenuMap[e.KContext.Screen.Value] = menu

		tt, _ := opentype.Parse(fonts.MPlus1pRegular_ttf)
		menu_face, _ := opentype.NewFace(tt, &opentype.FaceOptions{
			Size:    menu.MenuLayout.DPI,
			DPI:     menu.MenuLayout.Size,
			Hinting: font.HintingVertical,
		})
		for i, menu_item := range menu.MenuItems {
			if menu.CurrentSelected == i {
				text.Draw(screen, menu_item.Label+"\n", menu_face, 8, 24*(i+1), color.RGBA{R: 0, G: 255, B: 0, A: 100})
			} else {
				text.Draw(screen, menu_item.Label+"\n", menu_face, 8, 24*(i+1), color.White)
			}
		}
		if e.Controller.Action3Pressed() {
			os.Exit(1)
		}
	}
	e.Controller.Clear()
}
