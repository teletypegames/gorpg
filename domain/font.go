package domain

import (
	"game/konstructor"
	"image/color"
)

func GetDefaultFontLayout() konstructor.FontLayout {
	return konstructor.FontLayout{
		Path:          "assets/fonts/ArcadeClassic.ttf",
		DPI:           72,
		Size:          24,
		Color:         color.White,
		SelectedColor: color.RGBA{R: 255, G: 0, B: 0, A: 100},
	}
}
