package konstructor

type KContext struct {
	ScreenType        ScreenType
	ActiveMenu        MenuMapKey
	ActiveDialog      DialogMapKey
	CurrentLevel      int
	CurrentPlayground int
	Multiplayer       bool
}

func (c *KContext) ScreenTypeIs(name ScreenType) bool {
	return c.ScreenType == name
}
