package entity

import (
	"bytes"
	"image"
	"io/ioutil"
	"log"
)

type Render struct {
	Image   string
	Width   int
	Height  int
	Visible bool
}

func (ro *Render) GetImage() image.Image {
	file, err := ioutil.ReadFile(ro.Image)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(bytes.NewReader(file))
	if err != nil {
		log.Fatal(err)
	}
	return img
}
