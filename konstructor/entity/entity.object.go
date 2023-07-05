package entity

type ObjectType struct {
	ID            string
	RenderOptions RenderOptions
}

type Object struct {
	ID       string
	Type     ObjectType
	Position Position
}
