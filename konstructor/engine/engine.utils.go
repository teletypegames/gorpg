package engine

import (
	"bytes"
	"game/konstructor/entity"
	"image"
	"io/ioutil"
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func LoadImage(path string) *ebiten.Image {

	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(bytes.NewReader(file))
	if err != nil {
		log.Fatal(err)
	}

	out, err := ebiten.NewImageFromImage(img, 0)
	if err != nil {
		log.Fatal(err)
	}

	return out
}

func GetFontFace(layout entity.FontLayout) font.Face {
	file, _ := ioutil.ReadFile(layout.Path)
	tt, _ := opentype.Parse(file)
	face, _ := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    layout.DPI,
		DPI:     layout.Size,
		Hinting: font.HintingVertical,
	})
	return face
}
