package konstructor

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

type Config struct {
	Name   string
	Screen *ScreenConfig
	KeyMap KeyMap
}
