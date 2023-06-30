package constructor

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

type GameSettings struct {
	Name           string
	ScreenSettings ScreenSettings
}

type Constructor struct {
	Context      ContextInterface
	Domain       DomainInterface
	Controller   Controller
	GameSettings GameSettings
}

func (c Constructor) Init() {
	ebiten.SetWindowSize(c.GameSettings.ScreenSettings.Width, c.GameSettings.ScreenSettings.Height)
	ebiten.SetWindowTitle(c.GameSettings.Name)
	if err := ebiten.RunGame(&GameEngine{
		Constructor: c,
	}); err != nil {
		log.Fatal(err)
	}
}
