package domain

import (
	"game/src/konstructor"
)

const (
	DialogTest konstructor.DialogMapKey = "DialogTest"
)

func (d *Domain) InitDialog() {
	d.DialogMap = konstructor.DialogMap{
		DialogTest: {
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

func (d *Domain) GetDialog(name konstructor.DialogMapKey) konstructor.Dialog {
	value, _ := d.DialogMap[name]
	return value
}

func (d *Domain) SetDialog(name konstructor.DialogMapKey, dialog konstructor.Dialog) {
	d.DialogMap[name] = dialog
}
