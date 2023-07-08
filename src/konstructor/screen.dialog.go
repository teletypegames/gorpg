package konstructor

import "image/color"

type DialogMapKey string

type DialogMap map[DialogMapKey]Dialog

// Layout of dialog
type DialogLayout struct {
	Render     Render
	ChoiceFont FontLayout
}

// Choicable answer in a dialog (by player)
type DialogChoice struct {
	ID      string
	Label   string
	Handler func()
}

// Dialog between player and NPC
type Dialog struct {
	CurrentSelected int
	Layout          DialogLayout
	Choices         []DialogChoice
}

// Return with color of choice
func (dialog *Dialog) GetChoiceColor(i int) color.Color {
	if dialog.CurrentSelected == i {
		return dialog.Layout.ChoiceFont.SelectedColor
	} else {
		return dialog.Layout.ChoiceFont.Color
	}
}
