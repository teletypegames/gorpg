package konstructor

import (
	"game/konstructor/entity"

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
	k.SetWindow()
	k.Run()
}

func (k Konstructor) SetWindow() {
	ebiten.SetWindowSize(k.Settings.Screen.Width, k.Settings.Screen.Height)
	ebiten.SetWindowTitle(k.Settings.Name)
}

func (k Konstructor) Run() {
	ebiten.RunGame(&Engine{
		KContext:   k.KContext,
		Domain:     k.Domain,
		Controller: k.Controller,
		Settings:   k.Settings,
	})
}
