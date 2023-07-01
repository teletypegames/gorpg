package domain

type LevelSection struct {
}

type Level struct {
	LevelSections []LevelSection
}

type LevelManager struct {
}

func (lm LevelManager) Create(screens []LevelSection) Level {
	return Level{
		LevelSections: screens,
	}
}

func (lm LevelManager) List() []Level {
	return []Level{
		lm.Create(LevelOne()),
	}
}

func (lm LevelManager) Get() Level {
	return Level{}
}
