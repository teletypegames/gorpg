package domain

type Screen struct {
}

type ScreenManager struct {
	Context Context
}

func (sm ScreenManager) GetCurrent() Screen {
	return Screen{}
}
