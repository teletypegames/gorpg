package konstructor

type NPCTypeMap map[string]NPCType

type NPCType struct {
	ID     string
	Render Render
}

type NPC struct {
	ID       string
	Type     NPCType
	Position Position
}
