package entity

type KContext struct {
	ScreenType   ScreenType
	ScreenValue  string
	CurrentLevel int
	Players      []Player
}

func (c *KContext) getPrimaryPlayer() Player {
	return c.Players[0]
}

func (c *KContext) ScreenTypeIs(name ScreenType) bool {
	return c.ScreenType == name
}
