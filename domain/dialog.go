package domain

import (
	"game/konstructor/entity"
	"image/color"
)

func (d *Domain) InitDialog() {
	d.DialogMap = entity.DialogMap{
		"TestDialog": {
			Layout: entity.DialogLayout{
				ChoiceFont: entity.FontLayout{
					DPI:           72,
					Size:          24,
					Color:         color.White,
					SelectedColor: color.RGBA{R: 0, G: 255, B: 0, A: 100},
				},
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
