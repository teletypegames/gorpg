package entity

type ScreenSettings struct {
	Width  int
	Height int
}

type Settings struct {
	Name   string
	Screen *ScreenSettings
}
