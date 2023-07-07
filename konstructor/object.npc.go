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

func GetNPCTypeImagePath(name NPCTypeMapKey) string {
	return GetObjectDirectory(NPCObjectType) + string(name) + ".png"
}
