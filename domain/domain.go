package domain

import (
	"game/konstructor/entity"
)

type Domain struct {
	Context   Context
	MenuMap   entity.MenuMap
	DialogMap entity.DialogMap
	Levels    []entity.Level
}

func (d *Domain) Init() {
	d.InitMenu()
	d.InitDialog()
	d.InitLevel()
}
