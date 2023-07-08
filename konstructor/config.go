package konstructor

import "image/color"

type ResolutionSet struct {
	Width  int
	Height int
}

var QVGAResolutionSet = ResolutionSet{
	Width:  320,
	Height: 240,
}

var VGAResolutionSet = ResolutionSet{
	Width:  640,
	Height: 480,
}

var SVGAResolutionSet = ResolutionSet{
	Width:  800,
	Height: 600,
}

var XGAResolutionSet = ResolutionSet{
	Width:  1024,
	Height: 768,
}

var ResolutionSets = []ResolutionSet{
	QVGAResolutionSet,
	VGAResolutionSet,
	SVGAResolutionSet,
	XGAResolutionSet,
}

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

func (c *Config) SetResolution(set ResolutionSet) {
	c.Screen.Width = set.Width
	c.Screen.Height = set.Height
}
