package main

import (
	"game/src/domain"
	"game/src/konstructor"
	"game/src/konstructor/engines/easy_ebitengine"
)

func main() {
	k := konstructor.Konstructor{
		EngineWrapper: &easy_ebitengine.EngineWrapper{},
		Domain:        &domain.Domain{},
	}
	k.Init()
}
