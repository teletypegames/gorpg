package easy_ebitengine

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

func (e *Engine) DialogUpdate() {
	dialog := e.Domain.GetDialog(e.KContext.ScreenValue)

	if e.UpPressed() && dialog.CurrentSelected != 0 {
		dialog.CurrentSelected--
	}

	if e.DownPressed() && dialog.CurrentSelected != len(dialog.Choices)-1 {
		dialog.CurrentSelected++
	}

	if e.Action0Pressed() {
		dialog.Choices[dialog.CurrentSelected].Handler()
	}
	e.Domain.SetDialog(e.KContext.ScreenValue, dialog)
}

func (e *Engine) DialogDraw(screen *ebiten.Image) {
	dialog := e.Domain.GetDialog(e.KContext.ScreenValue)

	face := dialog.Layout.ChoiceFont.GetFontFace()

	for i, choice := range dialog.Choices {
		offset := int(dialog.Layout.ChoiceFont.Size) * (i + 1)
		text.Draw(screen, choice.Label+"\n", face, 8, offset, dialog.GetChoiceColor(i))
	}
}
