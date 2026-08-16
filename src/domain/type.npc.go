package domain

import "git.teletypegames.org/engines/gorpg/src/konstructor"

const (
	TestNPCType konstructor.NPCTypeMapKey = "test_npc"
)

func (d *Domain) InitNPCType() {
	d.NPCTypeMap = konstructor.NPCTypeMap{
		TestNPCType: {
			Render: konstructor.Render{
				Image: konstructor.GetNPCTypeImagePath(TestNPCType),
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
