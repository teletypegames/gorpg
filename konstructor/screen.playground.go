package konstructor

type Position struct {
	X int
	Y int
	Z int
}

type Playground struct {
	Render    Render
	Platforms []Platform
	NPCs      []NPC
	Items     []Item
}

type Level struct {
	ID          string
	Name        string
	Playgrounds []Playground
}
