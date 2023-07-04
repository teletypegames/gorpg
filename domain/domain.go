package domain

import (
	"game/konstructor"
)

type Domain struct {
	Context   Context
	MenuMap   konstructor.MenuMap
	DialogMap konstructor.DialogMap
}

func (d *Domain) Init() {
	d.InitMenu()
	d.InitDialog()
}
