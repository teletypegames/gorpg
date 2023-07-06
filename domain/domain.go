package domain

import (
	"game/konstructor"
)

type Domain struct {
	konstructor.DomainBaseFields
	Context Context
}

func (d *Domain) Init() {
	d.InitObjectType()
	d.InitMenu()
	d.InitDialog()
	d.InitLevel()
}
