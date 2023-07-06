package konstructor

import "image/color"

type DialogMap map[string]Dialog

type DialogLayout struct {
	Render     Render
	ChoiceFont FontLayout
}

type DialogChoice struct {
	ID      string
	Label   string
	Handler func()
}

type Dialog struct {
	CurrentSelected int
	Layout          DialogLayout
	Choices         []DialogChoice
}

func (dialog *Dialog) GetChoiceColor(i int) color.Color {
	if dialog.CurrentSelected == i {
		return dialog.Layout.ChoiceFont.SelectedColor
	} else {
		return dialog.Layout.ChoiceFont.Color
	}
}
