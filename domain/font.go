package domain

import (
	"game/konstructor"
	"image/color"
)

const (
	DefaultFontName konstructor.FontName = "ArcadeClassic"
)

const (
	DefaultDPI = 72
)

func GetDefaultFontLayout() konstructor.FontLayout {
	return konstructor.FontLayout{
		Path:          konstructor.GetFontPath(DefaultFontName),
		DPI:           DefaultDPI,
		Size:          24,
		Color:         color.White,
		SelectedColor: color.RGBA{R: 255, G: 0, B: 0, A: 100},
	}
}

func GetHeaderFontLayout() konstructor.FontLayout {
	return konstructor.FontLayout{
		Path:          konstructor.GetFontPath(DefaultFontName),
		DPI:           DefaultDPI,
		Size:          12,
		Color:         color.White,
		SelectedColor: color.RGBA{R: 255, G: 0, B: 0, A: 100},
	}
}
