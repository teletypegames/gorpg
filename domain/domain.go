package domain

import (
	"game/konstructor"
)

type Domain struct {
	Context       Context
	MenuMap       konstructor.MenuMap
	DialogMap     konstructor.DialogMap
	Levels        []konstructor.Level
	ObjectTypeMap konstructor.ObjectTypeMap
}

func (d *Domain) Init() {
	d.InitMenu()
	d.InitDialog()
	d.InitLevel()
}
