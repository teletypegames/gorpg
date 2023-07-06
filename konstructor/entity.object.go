package konstructor

type ObjectTypeMap map[string]ObjectType

type ObjectType struct {
	ID     string
	Render Render
}

type Object struct {
	ID       string
	Type     ObjectType
	Position Position
}
