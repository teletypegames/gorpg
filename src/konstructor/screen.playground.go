package konstructor

type PlaygroundID string

type Position struct {
	X int
	Y int
	Z int
}

type Playground struct {
	ID        PlaygroundID
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

func GetPlaygroundImagePath(name PlaygroundID) string {
	return "assets/images/playgrounds/" + string(name) + ".png"
}
