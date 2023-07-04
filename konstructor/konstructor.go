package konstructor

import (
	"game/konstructor/engine"
	"game/konstructor/entity"

	"github.com/hajimehoshi/ebiten"
)

type Konstructor struct {
	Domain     entity.DomainInterface
	Controller *engine.Controller
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
	ebiten.RunGame(&engine.Engine{
		KContext:   k.KContext,
		Domain:     k.Domain,
		Controller: k.Controller,
		Settings:   k.Settings,
	})
}
