package domain

type ControllerInterface interface {
}

type PresenterInterface interface {
}

type Domain struct {
	Context    Context
	Controller ControllerInterface
	Presenter  PresenterInterface
}

func (d Domain) Init() {
	menu_manager := MenuManager{}
	menu_manager.Init()
	// level_manager := LevelManager{}
	screen_manager := ScreenManager{}
	screen_manager.GetCurrent()

}

func (d Domain) Start() {

}
