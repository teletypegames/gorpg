package domain

import "game/konstructor"

const (
	TestNPCType konstructor.NPCTypeMapKey = "test"
)

func (d *Domain) InitNPCType() {
	d.NPCTypeMap = konstructor.NPCTypeMap{
		TestNPCType: {
			Render: konstructor.Render{
				Image: "assets/images/npcs/test_npc.png",
			},
		},
	}
}

func (d *Domain) GetNPCTypeMap() konstructor.NPCTypeMap {
	return d.NPCTypeMap
}

func (d *Domain) GetNPCType(name konstructor.NPCTypeMapKey) konstructor.NPCType {
	value, _ := d.NPCTypeMap[name]
	return value
}
