package konstructor

type ObjectTypeMapKey string

type ObjectTypeMap map[ObjectTypeMapKey]ObjectType

type ObjectType struct {
	ID     string
	Render Render
}

type Object struct {
	ID       string
	Type     ObjectType
	Position Position
}
