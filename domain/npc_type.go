package domain

import "game/konstructor"

func (d *Domain) InitNPCType() {
	d.NPCTypeMap = konstructor.NPCTypeMap{}
}

func (d *Domain) GetNPCTypeMap() konstructor.NPCTypeMap {
	return d.NPCTypeMap
}

func (d *Domain) GetNPCType(name string) konstructor.NPCType {
	value, _ := d.NPCTypeMap[name]
	return value
}
