package konstructor

type ItemTypeMapKey string

type ItemTypeMap map[ItemTypeMapKey]ItemType

type ItemType struct {
	ID     string
	Render Render
}

type Item struct {
	ID       string
	Type     ItemType
	Position Position
}

func GetItemTypeImagePath(name ItemTypeMapKey) string {
	return GetObjectDirectory(ItemObjectType) + string(name) + ".png"
}
