package konstructor

type PlaygroundObject struct {
	Image string
	PosX  int
	PosY  int
}

type Playground struct {
	Background string
	Objects    []PlaygroundObject
}
