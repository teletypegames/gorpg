package engine

import (
	"fmt"
	"reflect"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func (e *Engine) WatchKeyPress() {
	values := reflect.ValueOf(e.Settings.KeyMap)

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
	return e.PressedKey == e.Settings.KeyMap.Up.(ebiten.Key)
}

func (e *Engine) DownPressed() bool {
	return e.PressedKey == e.Settings.KeyMap.Down.(ebiten.Key)
}

func (e *Engine) RightPressed() bool {
	return e.PressedKey == e.Settings.KeyMap.Right.(ebiten.Key)
}

func (e *Engine) LeftPressed() bool {
	return e.PressedKey == e.Settings.KeyMap.Left.(ebiten.Key)
}

func (e *Engine) Action0Pressed() bool {
	return e.PressedKey == e.Settings.KeyMap.Action0.(ebiten.Key)
}

func (e *Engine) Action1Pressed() bool {
	return e.PressedKey == e.Settings.KeyMap.Action1.(ebiten.Key)
}

func (e *Engine) Action2Pressed() bool {
	return e.PressedKey == e.Settings.KeyMap.Action2.(ebiten.Key)
}

func (e *Engine) Action3Pressed() bool {
	return e.PressedKey == e.Settings.KeyMap.Action3.(ebiten.Key)
}
