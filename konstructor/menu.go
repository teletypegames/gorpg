package konstructor

type MenuMap map[string]Menu

type MenuLayout struct {
	Background string
}

type MenuItem struct {
	ID      string
	Label   string
	Handler func()
}

type Menu struct {
	MenuLayout MenuLayout
	MenuItems  []MenuItem
}
