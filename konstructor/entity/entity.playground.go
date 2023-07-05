package entity

type Position struct {
	X int
	Y int
	Z int
}

type Playground struct {
	Render  Render
	Objects []Object
	NPCs    []NPC
	Items   []Item
}

type Level struct {
	ID          string
	Name        string
	Playgrounds []Playground
}
