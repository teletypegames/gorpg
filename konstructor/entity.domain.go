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
	GetMenu(name string) Menu
	SetMenu(name string, menu Menu)

	GetDialogMap() DialogMap
	GetDialog(name string) Dialog
	SetDialog(name string, menu Dialog)

	GetLevels() []Level
	GetLevel(index int) Level

	GetObjectTypeMap() ObjectTypeMap
	GetObjectType(name string) ObjectType

	GetItemTypeMap() ItemTypeMap
	GetItemType(name string) ItemType

	GetNPCTypeMap() NPCTypeMap
	GetNPCType(name string) NPCType

	AddToInventory(item *Item) bool
	RemoveFromInventory(item *Item) bool
	UseInventoryItem(item *Item) bool

	Process(DomainProcessArgs)
}

type DomainProcessArgs struct {
	Level    *Level
	KContext *KContext
}
