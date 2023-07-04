package konstructor

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

type ContextInterface interface {
}

type DomainInterface interface {
	Init()
	GetMenu(name string) Menu
	SetMenu(name string, menu Menu)
	GetDialog() Dialog
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
	KContext   *KContext
}

func (k Konstructor) Init() {
	k.Domain.Init()
	ebiten.SetWindowSize(k.Settings.Screen.Width, k.Settings.Screen.Height)
	ebiten.SetWindowTitle(k.Settings.Name)
	if err := ebiten.RunGame(&Engine{
		KContext:   k.KContext,
		Domain:     k.Domain,
		Controller: k.Controller,
		Settings:   k.Settings,
	}); err != nil {
		log.Fatal(err)
	}
}
