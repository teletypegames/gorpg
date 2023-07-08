package konstructor

type PlayerTypeID string

type InventoryItem struct {
	Item   Item
	Used   bool
	Active bool
}

// Player inventory
type Inventory struct {
	Items []InventoryItem
}

// Player type
type PlayerType struct {
	ID     PlayerTypeID
	Render Render
}

// Playable character
type Player struct {
	ID        string
	Type      PlayerType
	Position  Position
	Inventory Inventory
}

// Helper, which return with the path of a specific player type
func GetPlayerTypeImagePath(name PlayerTypeID) string {
	return GetObjectDirectory(PlayerObjectType) + string(name) + ".png"
}
