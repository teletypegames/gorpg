package domain

import (
	"game/konstructor"
)

func (d *Domain) InitDialog() {
	d.DialogMap = konstructor.DialogMap{
		"TestDialog": {
			Layout: konstructor.DialogLayout{
				ChoiceFont: GetDefaultFontLayout(),
			},
			Choices: []konstructor.DialogChoice{
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

func (d *Domain) GetDialogMap() konstructor.DialogMap {
	return d.DialogMap
}

func (d *Domain) GetDialog(name string) konstructor.Dialog {
	value, _ := d.DialogMap[name]
	return value
}

func (d *Domain) SetDialog(name string, dialog konstructor.Dialog) {
	d.DialogMap[name] = dialog
}
