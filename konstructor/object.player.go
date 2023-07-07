package konstructor

type PlayerTypeID string

type InventoryItem struct {
	Item   Item
	Used   bool
	Active bool
}

type Inventory struct {
	Items []InventoryItem
}

type PlayerType struct {
	ID     PlayerTypeID
	Render Render
}

type Player struct {
	ID        string
	Type      PlayerType
	Position  Position
	Inventory Inventory
}

func GetPlayerTypeImagePath(name PlayerTypeID) string {
	return GetObjectDirectory(PlayerObjectType) + string(name) + ".png"
}
