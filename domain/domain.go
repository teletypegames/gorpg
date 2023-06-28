package domain

type ControllerInterface interface {
	Init()
}

type PresenterInterface interface {
	Init()
}

type Domain struct {
	Context    Context
	Controller ControllerInterface
	Presenter  PresenterInterface
}

func (d Domain) Init() {
	d.Controller.Init()
	d.Presenter.Init()
	menu_manager := MenuManager{}
	menu_manager.Init()
	// level_manager := LevelManager{}
	screen_manager := ScreenManager{}
	screen_manager.GetCurrent()
}
