package konstructor

// Type of screen type for method input validation
type ScreenType string

const (
	MenuScreenType       ScreenType = "menu"
	DialogScreenType     ScreenType = "dialog"
	PlaygroundScreenType ScreenType = "playground"
	InventoryScreenType  ScreenType = "inventory"
)
