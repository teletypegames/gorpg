package domain

type User struct {
	Name string
}

type PlayerState struct {
	Name string
}

type Context struct {
	PlayerState PlayerState
	User        User
	Level       int
}
