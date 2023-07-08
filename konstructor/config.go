package konstructor

import "image/color"

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

type ScreenConfig struct {
	Width  int
	Height int
	Scale  int
}
type HeaderConfig struct {
	BackgroundColor color.Color
	Height          int
	FontLayout      FontLayout
}

type Config struct {
	Name   string
	Screen *ScreenConfig
	Header *HeaderConfig
	KeyMap KeyMap
}
