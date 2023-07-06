package entity

type KeyMap struct {
	Up      any
	Down    any
	Right   any
	Left    any
	Action0 any
	Action1 any
	Action2 any
	Action3 any
}

type ScreenSettings struct {
	Width  int
	Height int
	Scale  int
}

type Settings struct {
	Name   string
	Screen *ScreenSettings
	KeyMap KeyMap
}
