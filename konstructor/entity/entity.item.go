package entity

type ItemType struct {
	ID            string
	RenderOptions RenderOptions
}

type Item struct {
	ID       string
	Type     ItemType
	Position Position
}
