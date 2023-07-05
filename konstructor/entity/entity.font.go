package entity

import (
	"image/color"
	"io/ioutil"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type FontLayout struct {
	Path          string
	DPI           float64
	Size          float64
	Color         color.Color
	SelectedColor color.Color
}

func (fl *FontLayout) GetFontFace() font.Face {
	file, _ := ioutil.ReadFile(fl.Path)
	tt, _ := opentype.Parse(file)
	face, _ := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fl.DPI,
		DPI:     fl.Size,
		Hinting: font.HintingVertical,
	})
	return face
}
