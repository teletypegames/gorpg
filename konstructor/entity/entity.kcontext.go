package entity

type KContext struct {
	ScreenType   string
	ScreenValue  string
	CurrentLevel int
	Players      []Player
}

func (c *KContext) getPrimaryPlayer() Player {
	return c.Players[0]
}

func (c *KContext) ScreenTypeIs(name string) bool {
	return c.ScreenType == name
}
