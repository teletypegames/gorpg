package domain

import "game/konstructor"

func (d *Domain) InitItemType() {
	d.ItemTypeMap = konstructor.ItemTypeMap{}
}

func (d *Domain) GetItemTypeMap() konstructor.ItemTypeMap {
	return d.ItemTypeMap
}

func (d *Domain) GetItemType(name string) konstructor.ItemType {
	value, _ := d.ItemTypeMap[name]
	return value
}
