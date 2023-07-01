package domain

type User struct {
	Name string
}

type Player struct {
	Name string
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
