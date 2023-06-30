package konstructor

import (
	"fmt"
	"reflect"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type KeyMap struct {
	Up      ebiten.Key
	Down    ebiten.Key
	Right   ebiten.Key
	Left    ebiten.Key
	Action0 ebiten.Key
	Action1 ebiten.Key
	Action2 ebiten.Key
	Action3 ebiten.Key
}

type Controller struct {
	PressedKey ebiten.Key
	KeyMap     KeyMap
}

func (k *Controller) Watch() {
	values := reflect.ValueOf(k.KeyMap)

	for i := 0; i < values.NumField(); i++ {
		key := values.Field(i).Interface().(ebiten.Key)
		if inpututil.IsKeyJustPressed(key) {
			fmt.Printf("Key pressed: %s\n", key)
			k.PressedKey = key
		}
	}
}

func (k *Controller) UpPressed() bool {
	return k.PressedKey == k.KeyMap.Up
}

func (k *Controller) DownPressed() bool {
	return k.PressedKey == k.KeyMap.Down
}

func (k *Controller) RightPressed() bool {
	return k.PressedKey == k.KeyMap.Right
}

func (k *Controller) LeftPressed() bool {
	return k.PressedKey == k.KeyMap.Left
}

func (k *Controller) Action0Pressed() bool {
	return k.PressedKey == k.KeyMap.Action0
}

func (k *Controller) Action1Pressed() bool {
	return k.PressedKey == k.KeyMap.Action1
}

func (k *Controller) Action2Pressed() bool {
	return k.PressedKey == k.KeyMap.Action2
}

func (k *Controller) Action3Pressed() bool {
	return k.PressedKey == k.KeyMap.Action3
}
