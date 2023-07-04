package entity

import "image/color"

type FontLayout struct {
	DPI           float64
	Size          float64
	Color         color.Color
	SelectedColor color.Color
}
