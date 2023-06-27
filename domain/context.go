package domain

type User struct {
	name string
}

type Player struct {
	name string
}

type GameField struct {
	Level        Level
	LevelSection LevelSection
}

type Context struct {
	Player    Player
	User      User
	GameField GameField
}

func CreateContext() Context {
	return Context{}
}
