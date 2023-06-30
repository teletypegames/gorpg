package konstructor

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

type ContextInterface interface {
}

type DomainInterface interface {
	Init()
}

type ScreenSettings struct {
	Width  int
	Height int
}

type Settings struct {
	Name   string
	Screen ScreenSettings
}

type Konstructor struct {
	Context    ContextInterface
	Domain     DomainInterface
	Controller Controller
	Settings   Settings
}

func (k Konstructor) Init() {
	ebiten.SetWindowSize(k.Settings.Screen.Width, k.Settings.Screen.Height)
	ebiten.SetWindowTitle(k.Settings.Name)
	if err := ebiten.RunGame(&GameEngine{
		Konstructor: k,
	}); err != nil {
		log.Fatal(err)
	}
}
