package domain

import "git.teletypegames.org/engines/gorpg/src/konstructor"

const (
	SwordItemType konstructor.ItemTypeMapKey = "sword"
)

func (d *Domain) InitItemType() {
	d.ItemTypeMap = konstructor.ItemTypeMap{
		SwordItemType: {
			ID: "sword",
			Render: konstructor.Render{
				Image: konstructor.GetItemTypeImagePath(SwordItemType),
			},
		},
	}
}

func (d *Domain) GetItemTypeMap() konstructor.ItemTypeMap {
	return d.ItemTypeMap
}

func (d *Domain) GetItemType(name konstructor.ItemTypeMapKey) konstructor.ItemType {
	value, _ := d.ItemTypeMap[name]
	return value
}
