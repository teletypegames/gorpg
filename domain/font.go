package domain

import (
	"game/konstructor/entity"
	"image/color"
)

func GetDefaultFontLayout() entity.FontLayout {
	return entity.FontLayout{
		Path:          "fonts/ArcadeClassic.ttf",
		DPI:           72,
		Size:          24,
		Color:         color.White,
		SelectedColor: color.RGBA{R: 255, G: 0, B: 0, A: 100},
	}
}
