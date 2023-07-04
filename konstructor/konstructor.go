package konstructor

import (
	"game/konstructor/entity"
	"log"

	"github.com/hajimehoshi/ebiten"
)

type ContextInterface interface {
}

type DomainInterface interface {
	Init()
	GetMenu(name string) entity.Menu
	SetMenu(name string, menu entity.Menu)
	GetDialog(name string) entity.Dialog
	SetDialog(name string, menu entity.Dialog)
	GetLevel(index int) entity.Level
}

type Konstructor struct {
	Domain     DomainInterface
	Controller *Controller
	Settings   *entity.Settings
	KContext   *entity.KContext
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
