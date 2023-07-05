package entity

type ItemType struct {
	ID     string
	Render Render
}

type Item struct {
	ID       string
	Type     ItemType
	Position Position
}
