package konstructor

type ObjectType struct {
	ID     string
	Render Render
}

type Object struct {
	ID       string
	Type     ObjectType
	Position Position
}
