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

type Position struct {
	X int
	Y int
}

type Context struct {
	Player           Player
	User             User
	CurrentGameField GameField
	CurrentPosition  Position
}

func CreateContext() Context {
	return Context{
		Player: Player{},
		User:   User{},
	}
}
