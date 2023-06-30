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

type GameSettings struct {
	Name           string
	ScreenSettings ScreenSettings
}

type Konstructor struct {
	Context      ContextInterface
	Domain       DomainInterface
	Controller   Controller
	GameSettings GameSettings
}

func (k Konstructor) Init() {
	ebiten.SetWindowSize(k.GameSettings.ScreenSettings.Width, k.GameSettings.ScreenSettings.Height)
	ebiten.SetWindowTitle(k.GameSettings.Name)
	if err := ebiten.RunGame(&GameEngine{
		Konstructor: k,
	}); err != nil {
		log.Fatal(err)
	}
}
