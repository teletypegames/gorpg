package entity

type NPCType struct {
	ID            string
	RenderOptions RenderOptions
}

type NPC struct {
	ID       string
	Type     NPCType
	Position Position
}
