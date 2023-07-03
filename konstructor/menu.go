package konstructor

type MenuMap map[string]Menu

type MenuLayout struct {
	Background string
	DPI        float64
	Size       float64
}

type MenuItem struct {
	ID      string
	Label   string
	Handler func()
}

type Menu struct {
	CurrentSelected int
	MenuLayout      MenuLayout
	MenuItems       []MenuItem
}
