package entity

import "github.com/hajimehoshi/ebiten"

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

type ScreenSettings struct {
	Width  int
	Height int
}

type Settings struct {
	Name   string
	Screen *ScreenSettings
	KeyMap *KeyMap
}
