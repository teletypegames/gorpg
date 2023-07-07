package konstructor

type ObjectType string

const (
	ItemObjectType     ObjectType = "item"
	NPCObjectType      ObjectType = "npc"
	PlatformObjectType ObjectType = "platform"
)

type ObjectDirectoryMap map[ObjectType]string

var ObjectDirectories = ObjectDirectoryMap{
	ItemObjectType:     "items",
	NPCObjectType:      "npcs",
	PlatformObjectType: "platforms",
}

func GetObjectDirectory(name ObjectType) string {
	return "assets/images/" + ObjectDirectories[name] + "/"
}
