package entity

type InventoryItem struct {
	Item   Item
	Used   bool
	Active bool
}

type Inventory struct {
	Items []InventoryItem
}

type Player struct {
	ID        string
	Render    Render
	Position  Position
	Inventory Inventory
}
