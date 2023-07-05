package entity

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

	AddToInventory(item *Item) bool
	RemoveFromInventory(item *Item) bool
	UseInventoryItem(item *Item) bool

	Process(DomainProcessArgs)
}

type DomainProcessArgs struct {
	Level    *Level
	KContext *KContext
}
