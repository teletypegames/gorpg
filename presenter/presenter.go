package presenter

import (
	"game/domain"
	"log"

	"github.com/hajimehoshi/ebiten"
)

type Presenter struct {
	Context domain.Context
}

func (p Presenter) Init() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Game")
	if err := ebiten.RunGame(&GameEngine{
		Context: p.Context,
		Keyboard: Keyboard{
			KeyMap: CreateDefaultKeyMap(),
		},
	}); err != nil {
		log.Fatal(err)
	}
}
