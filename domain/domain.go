package domain

type PresenterInterface interface {
	Init()
}

type Domain struct {
	Context   Context
	Presenter PresenterInterface
}

func (d Domain) Init() {
	d.Presenter.Init()
	menu_manager := MenuManager{}
	menu_manager.Init()
	// level_manager := LevelManager{}
	screen_manager := ScreenManager{}
	screen_manager.GetCurrent()
}
