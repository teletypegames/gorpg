package domain

type Domain struct {
	Context Context
}

func (d Domain) Init() {
	menu_manager := MenuManager{}
	menu_manager.Init()
	// level_manager := LevelManager{}
	screen_manager := ScreenManager{}
	screen_manager.GetCurrent()
}
