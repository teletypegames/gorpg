package constructor

import (
	"game/domain"
	"log"

	"github.com/hajimehoshi/ebiten"
)

type DomainInterface interface {
	Init()
}

type Constructor struct {
	Context  domain.Context
	Domain   DomainInterface
	Keyboard Keyboard
}

func (p Constructor) Init() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Game")
	if err := ebiten.RunGame(&GameEngine{
		Context:  p.Context,
		Keyboard: p.Keyboard,
	}); err != nil {
		log.Fatal(err)
	}
}
