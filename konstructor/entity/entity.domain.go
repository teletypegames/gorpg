package entity

type DomainInterface interface {
	Init()
	GetMenu(name string) Menu
	SetMenu(name string, menu Menu)
	GetDialog(name string) Dialog
	SetDialog(name string, menu Dialog)
	GetLevel(index int) Level
	AddToInventory(item *Item) bool
	RemoveFromInventory(item *Item) bool
	UseInventoryItem(item *Item) bool
}
