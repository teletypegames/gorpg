package engine

import (
	"fmt"
	"game/konstructor/entity"
	"reflect"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Controller struct {
	PressedKey ebiten.Key
	KeyMap     entity.KeyMap
}

func (c *Controller) Watch() {
	values := reflect.ValueOf(c.KeyMap)

	for i := 0; i < values.NumField(); i++ {
		key := values.Field(i).Interface().(ebiten.Key)
		if inpututil.IsKeyJustPressed(key) {
			fmt.Printf("Key pressed: %s\n", key)
			c.PressedKey = key
		}
	}
}

func (c *Controller) Clear() {
	c.PressedKey = Controller{}.PressedKey
}

func (c *Controller) UpPressed() bool {
	return c.PressedKey == c.KeyMap.Up
}

func (c *Controller) DownPressed() bool {
	return c.PressedKey == c.KeyMap.Down
}

func (c *Controller) RightPressed() bool {
	return c.PressedKey == c.KeyMap.Right
}

func (c *Controller) LeftPressed() bool {
	return c.PressedKey == c.KeyMap.Left
}

func (c *Controller) Action0Pressed() bool {
	return c.PressedKey == c.KeyMap.Action0
}

func (c *Controller) Action1Pressed() bool {
	return c.PressedKey == c.KeyMap.Action1
}

func (c *Controller) Action2Pressed() bool {
	return c.PressedKey == c.KeyMap.Action2
}

func (c *Controller) Action3Pressed() bool {
	return c.PressedKey == c.KeyMap.Action3
}
