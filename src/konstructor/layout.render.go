package konstructor

import (
	"bytes"
	"image"
	"io/ioutil"
	"log"
)

type Render struct {
	Image       string
	Width       int
	Height      int
	Visible     bool
	cachedImage image.Image
}

func (ro *Render) GetImage() image.Image {
	if ro.cachedImage != nil {
		return ro.cachedImage
	}

	file, err := ioutil.ReadFile(ro.Image)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(bytes.NewReader(file))
	if err != nil {
		log.Fatal(err)
	}
	ro.cachedImage = img
	return img
}
