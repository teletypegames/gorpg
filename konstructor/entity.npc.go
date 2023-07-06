package konstructor

type NPCType struct {
	ID     string
	Render Render
}

type NPC struct {
	ID       string
	Type     NPCType
	Position Position
}
