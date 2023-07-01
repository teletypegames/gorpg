package domain

import "game/konstructor"

type PlaygroundManager struct {
	Context Context
}

func (pm PlaygroundManager) GetCurrent() konstructor.Playground {
	return konstructor.Playground{}
}
