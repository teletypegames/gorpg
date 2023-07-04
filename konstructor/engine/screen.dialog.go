package engine

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

func (e *Engine) DialogDraw(screen *ebiten.Image) {
	dialog := e.Domain.GetDialog(e.KContext.Screen.Value)

	face := GetFontFace(dialog.Layout.ChoiceFont)

	for i, choice := range dialog.Choices {
		offset := int(dialog.Layout.ChoiceFont.Size) * (i + 1)
		text.Draw(screen, choice.Label+"\n", face, 8, offset, dialog.GetChoiceColor(i))
	}
}

func (e *Engine) DialogUpdate() {
	dialog := e.Domain.GetDialog(e.KContext.Screen.Value)

	if e.Controller.UpPressed() && dialog.CurrentSelected != 0 {
		dialog.CurrentSelected--
	}

	if e.Controller.DownPressed() && dialog.CurrentSelected != len(dialog.Choices)-1 {
		dialog.CurrentSelected++
	}

	if e.Controller.Action0Pressed() {
		dialog.Choices[dialog.CurrentSelected].Handler()
	}
	e.Domain.SetDialog(e.KContext.Screen.Value, dialog)
}
