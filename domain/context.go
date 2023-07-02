package domain

type User struct {
	Name string
}

type Player struct {
	Name string
}

type Context struct {
	Player Player
	User   User
}
