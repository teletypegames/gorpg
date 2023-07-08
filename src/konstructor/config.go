package konstructor

import "image/color"

// Windows resolution set
type ResolutionSet struct {
	Width  int
	Height int
}

// QVGA resolution set
var QVGAResolutionSet = ResolutionSet{
	Width:  320,
	Height: 240,
}

// VGA resolution set
var VGAResolutionSet = ResolutionSet{
	Width:  640,
	Height: 480,
}

// SVGA resolution set
var SVGAResolutionSet = ResolutionSet{
	Width:  800,
	Height: 600,
}

// XGA resolution set
var XGAResolutionSet = ResolutionSet{
	Width:  1024,
	Height: 768,
}

// List of resolution sets
var ResolutionSets = []ResolutionSet{
	QVGAResolutionSet,
	VGAResolutionSet,
	SVGAResolutionSet,
	XGAResolutionSet,
}

// Keyboard layout configuration
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

// Screen configuration
type ScreenConfig struct {
	Width  int
	Height int
	Scale  int
}

// Header configuration
type HeaderConfig struct {
	BackgroundColor color.Color
	Height          int
	FontLayout      FontLayout
}

// Game configuration
type Config struct {
	Name   string
	Screen *ScreenConfig
	Header *HeaderConfig
	KeyMap KeyMap
}

// Set screen resolution config to a resolution set
func (c *Config) SetResolution(set ResolutionSet) {
	c.Screen.Width = set.Width
	c.Screen.Height = set.Height
}
