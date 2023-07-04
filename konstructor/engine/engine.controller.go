package engine

import (
	"fmt"
	"reflect"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func (e *Engine) WatchKeyPress() {
	values := reflect.ValueOf(e.KeyMap)

	for i := 0; i < values.NumField(); i++ {
		key := values.Field(i).Interface().(ebiten.Key)
		if inpututil.IsKeyJustPressed(key) {
			fmt.Printf("Key pressed: %s\n", key)
			e.PressedKey = key
		}
	}
}

func (e *Engine) ClearKeyPresed() {
	e.PressedKey = Engine{}.PressedKey
}

func (e *Engine) UpPressed() bool {
	return e.PressedKey == e.KeyMap.Up
}

func (e *Engine) DownPressed() bool {
	return e.PressedKey == e.KeyMap.Down
}

func (e *Engine) RightPressed() bool {
	return e.PressedKey == e.KeyMap.Right
}

func (e *Engine) LeftPressed() bool {
	return e.PressedKey == e.KeyMap.Left
}

func (e *Engine) Action0Pressed() bool {
	return e.PressedKey == e.KeyMap.Action0
}

func (e *Engine) Action1Pressed() bool {
	return e.PressedKey == e.KeyMap.Action1
}

func (e *Engine) Action2Pressed() bool {
	return e.PressedKey == e.KeyMap.Action2
}

func (e *Engine) Action3Pressed() bool {
	return e.PressedKey == e.KeyMap.Action3
}
