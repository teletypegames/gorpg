package domain

import (
	"game/konstructor"
)

type Domain struct {
	konstructor.DomainBaseFields
	Context Context
}

func (d *Domain) Init() {
	d.InitPlatformType()
	d.InitItemType()
	d.InitNPCType()
	d.InitMenu()
	d.InitDialog()
	d.InitLevel()
}
