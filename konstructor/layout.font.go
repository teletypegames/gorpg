package konstructor

import (
	"image/color"
	"io/ioutil"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type FontName string

type FontLayout struct {
	Path           string
	DPI            float64
	Size           float64
	Color          color.Color
	SelectedColor  color.Color
	cachedFontFace font.Face
}

func (fl *FontLayout) GetFontFace(scale int) font.Face {
	if fl.cachedFontFace != nil {
		return fl.cachedFontFace
	}
	file, _ := ioutil.ReadFile(fl.Path)
	true_type, _ := opentype.Parse(file)
	font_face, _ := opentype.NewFace(true_type, &opentype.FaceOptions{
		Size:    fl.Size * float64(scale),
		DPI:     fl.DPI,
		Hinting: font.HintingVertical,
	})
	fl.cachedFontFace = font_face
	return font_face
}

func GetFontPath(name FontName) string {
	return "assets/fonts/" + string(name) + ".ttf"
}
