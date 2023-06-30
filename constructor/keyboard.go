package constructor

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

type Keyboard struct {
	PressedKey ebiten.Key
	KeyMap     KeyMap
}

func (k *Keyboard) Watch() {
	values := reflect.ValueOf(k.KeyMap)

	for i := 0; i < values.NumField(); i++ {
		key := values.Field(i).Interface().(ebiten.Key)
		if inpututil.IsKeyJustPressed(key) {
			fmt.Printf("Key pressed: %s\n", key)
			k.PressedKey = key
		}
	}
}

func (k *Keyboard) UpPressed() bool {
	return k.PressedKey == k.KeyMap.Up
}

func (k *Keyboard) DownPressed() bool {
	return k.PressedKey == k.KeyMap.Down
}

func (k *Keyboard) RightPressed() bool {
	return k.PressedKey == k.KeyMap.Right
}

func (k *Keyboard) LeftPressed() bool {
	return k.PressedKey == k.KeyMap.Left
}

func (k *Keyboard) Action0Pressed() bool {
	return k.PressedKey == k.KeyMap.Action0
}

func (k *Keyboard) Action1Pressed() bool {
	return k.PressedKey == k.KeyMap.Action1
}

func (k *Keyboard) Action2Pressed() bool {
	return k.PressedKey == k.KeyMap.Action2
}

func (k *Keyboard) Action3Pressed() bool {
	return k.PressedKey == k.KeyMap.Action3
}
