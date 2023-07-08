package konstructor

// Object type type for function input validation
type ObjectType string

const (
	ItemObjectType     ObjectType = "item"
	NPCObjectType      ObjectType = "npc"
	PlatformObjectType ObjectType = "platform"
	PlayerObjectType   ObjectType = "player"
)

// Map of object type's directories
type ObjectDirectoryMap map[ObjectType]string

var ObjectDirectories = ObjectDirectoryMap{
	ItemObjectType:     "items",
	NPCObjectType:      "npcs",
	PlatformObjectType: "platforms",
	PlayerObjectType:   "players",
}

// Return with the path of a specific object type asset directory
func GetObjectDirectory(name ObjectType) string {
	return "assets/images/" + ObjectDirectories[name] + "/"
}
