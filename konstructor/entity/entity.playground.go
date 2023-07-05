package entity

type Position struct {
	X int
	Y int
	Z int
}

type RenderOptions struct {
	Image   string
	Width   int
	Height  int
	Visible bool
}

type Playground struct {
	Background string
	Objects    []Object
	NPCs       []NPC
	Items      []Item
}

type Level struct {
	ID          string
	Name        string
	Playgrounds []Playground
}
