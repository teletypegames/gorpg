package domain

import (
	"git.teletypegames.org/engines/gorpg/src/konstructor"
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
	d.InitPlayer()
	d.InitLevel()
}
