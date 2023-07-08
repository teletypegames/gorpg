package konstructor

type ItemTypeMapKey string

type ItemTypeMap map[ItemTypeMapKey]ItemType

// Type of item
type ItemType struct {
	ID     string
	Render Render
}

// Playground usable item, which can be put in the Inventory
type Item struct {
	ID       string
	Type     ItemType
	Position Position
}

// Helper, which return with the path of a specific item type
func GetItemTypeImagePath(name ItemTypeMapKey) string {
	return GetObjectDirectory(ItemObjectType) + string(name) + ".png"
}
