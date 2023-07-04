package domain

import (
	"game/konstructor"
	"image/color"
)

type Domain struct {
	Context Context
	MenuMap konstructor.MenuMap
}

func (d *Domain) Init() {
	d.InitMenu()
}

func (d *Domain) GetDialog() konstructor.Dialog {
	return konstructor.Dialog{
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
	}
}
