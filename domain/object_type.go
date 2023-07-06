package domain

import "game/konstructor"

func (d *Domain) InitObjectType() {
	d.ObjectTypeMap = konstructor.ObjectTypeMap{
		"TestObjectType": konstructor.ObjectType{
			ID: "test_object_type",
			Render: konstructor.Render{
				Image:  "assets/images/objects/test_object.png",
				Width:  30,
				Height: 30,
			},
		},
	}
}

func (d *Domain) GetObjectTypeMap() konstructor.ObjectTypeMap {
	return d.ObjectTypeMap
}

func (d *Domain) GetObjectType(name konstructor.ObjectTypeMapKey) konstructor.ObjectType {
	value, _ := d.ObjectTypeMap[name]
	return value
}
