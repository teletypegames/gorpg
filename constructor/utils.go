package constructor

import (
	"bytes"
	"fmt"
	"image"
	"io/ioutil"
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten"
)

func LoadImage(path string) *ebiten.Image {

	file, err := ioutil.ReadFile(path) //read the content of file
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(bytes.NewReader(file))
	if err != nil {
		fmt.Println("decode error")
		log.Fatal(err)
	}

	out, err := ebiten.NewImageFromImage(img, 0)
	if err != nil {
		log.Fatal(err)
	}

	return out
}
