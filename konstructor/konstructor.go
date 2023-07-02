package konstructor

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

type ContextInterface interface {
}

type DomainInterface interface {
	GetMenuMap() MenuMap
	Init()
}

type ScreenSettings struct {
	Width  int
	Height int
}

type Settings struct {
	Name   string
	Screen *ScreenSettings
}

type Konstructor struct {
	Domain     DomainInterface
	Controller *Controller
	Settings   *Settings
}

func (k Konstructor) Init() {
	ebiten.SetWindowSize(k.Settings.Screen.Width, k.Settings.Screen.Height)
	ebiten.SetWindowTitle(k.Settings.Name)
	if err := ebiten.RunGame(&Engine{
		Logic: &Logic{
			Domain:     k.Domain,
			Controller: k.Controller,
			Settings:   k.Settings,
		},
	}); err != nil {
		log.Fatal(err)
	}
}
