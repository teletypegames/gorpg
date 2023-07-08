package konstructor

// Base fields for Domain layer.
// You need to include this of your game's domain struct
type DomainBaseFields struct {
	MenuMap         MenuMap
	DialogMap       DialogMap
	Levels          []Level
	Players         []Player
	PlatformTypeMap PlatformTypeMap
	ItemTypeMap     ItemTypeMap
	NPCTypeMap      NPCTypeMap
}

// Domain interface.
// You need to implement this interface in your game's domain struct
type DomainInterface interface {
	// Domain initialization. This variable will invoke only once when the game starts
	Init()

	// Return with the default KContext object
	GetInitialKContext() *KContext
	// Return with the default config object
	GetInitialConfig() *Config

	// Return with the menu map
	GetMenuMap() MenuMap
	// Return with a specific menu
	GetMenu(name MenuMapKey) Menu
	// Owerride a specific menu
	SetMenu(name MenuMapKey, menu Menu)

	// Return with the dialog map
	GetDialogMap() DialogMap
	// Return with a specific dialog
	GetDialog(name DialogMapKey) Dialog
	// Override a specific menu
	SetDialog(name DialogMapKey, menu Dialog)

	// Get levels
	GetLevels() []Level
	// Get a specific level
	GetLevel(index int) Level

	// Get players for multiplayer game
	GetPlayers() []Player
	// Get a specific player for multiplayer game
	GetPlayer(index int) Player
	// Get the player for simple player game
	GetDefaultPlayer() Player

	// Return with the platform type map
	GetPlatformTypeMap() PlatformTypeMap
	// Return with a specific platform type
	GetPlatformType(name PlatformTypeMapKey) PlatformType

	// Return with the item type map
	GetItemTypeMap() ItemTypeMap
	// Return with a specific item type
	GetItemType(name ItemTypeMapKey) ItemType

	// Return with the NPC type map
	GetNPCTypeMap() NPCTypeMap
	// Return with a specific NPC type
	GetNPCType(name NPCTypeMapKey) NPCType

	// Add item to inventory
	AddToInventory(item *Item) bool
	// Remove item from inventory
	RemoveFromInventory(item *Item) bool
	// Use inventory item
	UseInventoryItem(item *Item) bool

	// Process game flow by Domain layer
	Process(DomainProcessArgs)
}

type DomainProcessArgs struct {
	Level    *Level
	KContext *KContext
}
