package entity

type KContextScreen struct {
	Type  string
	Value string
}

type KContext struct {
	Screen       KContextScreen
	CurrentLevel int
}
