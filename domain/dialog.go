package domain

import (
	"game/konstructor/entity"
)

func (d *Domain) InitDialog() {
	d.DialogMap = entity.DialogMap{
		"TestDialog": {
			Layout: entity.DialogLayout{
				ChoiceFont: GetDefaultFontLayout(),
			},
			Choices: []entity.DialogChoice{
				{
					ID:    "one",
					Label: "One",
				},
				{
					ID:    "two",
					Label: "Two",
				},
			},
		},
	}
}

func (d *Domain) GetDialogMap() entity.DialogMap {
	return d.DialogMap
}

func (d *Domain) GetDialog(name string) entity.Dialog {
	value, _ := d.DialogMap[name]
	return value
}

func (d *Domain) SetDialog(name string, dialog entity.Dialog) {
	d.DialogMap[name] = dialog
}
