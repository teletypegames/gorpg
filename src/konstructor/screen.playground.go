package konstructor

// ID of playground for method input validation
type PlaygroundID string

// Position configuration
type Position struct {
	X int
	Y int
	Z int
}

// Displayable section of a level
type Playground struct {
	ID        PlaygroundID
	Render    Render
	Platforms []Platform
	NPCs      []NPC
	Items     []Item
}

// Game level
type Level struct {
	ID          string
	Name        string
	Playgrounds []Playground
}

// Return with the path of a specific playground background
func GetPlaygroundImagePath(name PlaygroundID) string {
	return "assets/images/playgrounds/" + string(name) + ".png"
}
