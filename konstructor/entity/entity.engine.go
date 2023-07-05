package entity

type EngineInterface interface {
	ClearKeyPresed()
	UpPressed() bool
	DownPressed() bool
	RightPressed() bool
	LeftPressed() bool
	Action0Pressed() bool
	Action1Pressed() bool
	Action2Pressed() bool
	Action3Pressed() bool
	ScreenTypeIs(name string) bool
	Update(screen any) error
	Draw(screen any)
	AddToInventory(item *Item)
	RemoveFromInventory(item *Item)
	UseInventoryItem(item *Item)
	DialogDraw(screen any)
	DialogUpdate()
	MenuDraw(screen any)
	MenuUpdate()
	PlaygroundDraw(screen any)
	PlaygroundUpdate()
}

type EngineOptions struct {
	Engine   EngineInterface
	Domain   DomainInterface
	KContext KContext
	Settings Settings
}

type EngineWrapperInterface interface {
	Init(options EngineOptions)
	Run()
}
