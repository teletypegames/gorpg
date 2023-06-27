package domain

type Screen struct {
	Controller ControllerInterface
}

type ScreenManager struct {
	Context Context
}

func (sm ScreenManager) GetCurrent() Screen {
	return Screen{}
}
