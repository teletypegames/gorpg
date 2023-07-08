package konstructor

// Framework context (state)
type KContext struct {
	ScreenType        ScreenType
	ActiveMenu        MenuMapKey
	ActiveDialog      DialogMapKey
	CurrentLevel      int
	CurrentPlayground int
	Multiplayer       bool
	LiveCount         int
}

// Check the current screen type
func (c *KContext) ScreenTypeIs(name ScreenType) bool {
	return c.ScreenType == name
}
