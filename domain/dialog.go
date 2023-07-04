package domain

import (
	"game/konstructor"
	"image/color"
)

func (d *Domain) InitDialog() {
	d.DialogMap = konstructor.DialogMap{
		"TestDialog": {
			Layout: konstructor.DialogLayout{
				ChoiceFont: konstructor.FontLayout{
					DPI:           72,
					Size:          24,
					Color:         color.White,
					SelectedColor: color.RGBA{R: 0, G: 255, B: 0, A: 100},
				},
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

func (d *Domain) GetDialog(name string) konstructor.Dialog {
	value, _ := d.DialogMap[name]
	return value
}

func (d *Domain) SetDialog(name string, dialog konstructor.Dialog) {
	d.DialogMap[name] = dialog
}
