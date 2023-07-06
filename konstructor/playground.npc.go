package konstructor

type NPCTypeMapKey string

type NPCTypeMap map[NPCTypeMapKey]NPCType

type NPCType struct {
	ID     string
	Render Render
}

type NPC struct {
	ID       string
	Type     NPCType
	Position Position
}
