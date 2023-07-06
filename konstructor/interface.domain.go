package konstructor

type DomainBaseFields struct {
	MenuMap       MenuMap
	DialogMap     DialogMap
	Levels        []Level
	ObjectTypeMap ObjectTypeMap
	ItemTypeMap   ItemTypeMap
	NPCTypeMap    NPCTypeMap
}

type DomainInterface interface {
	Init()

	GetMenuMap() MenuMap
	GetMenu(name MenuMapKey) Menu
	SetMenu(name MenuMapKey, menu Menu)

	GetDialogMap() DialogMap
	GetDialog(name DialogMapKey) Dialog
	SetDialog(name DialogMapKey, menu Dialog)

	GetLevels() []Level
	GetLevel(index int) Level

	GetObjectTypeMap() ObjectTypeMap
	GetObjectType(name ObjectTypeMapKey) ObjectType

	GetItemTypeMap() ItemTypeMap
	GetItemType(name ItemTypeMapKey) ItemType

	GetNPCTypeMap() NPCTypeMap
	GetNPCType(name NPCTypeMapKey) NPCType

	AddToInventory(item *Item) bool
	RemoveFromInventory(item *Item) bool
	UseInventoryItem(item *Item) bool

	Process(DomainProcessArgs)
}

type DomainProcessArgs struct {
	Level    *Level
	KContext *KContext
}
