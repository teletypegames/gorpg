package main

import (
	"game/domain"
	"game/konstructor"
	"game/konstructor/engines/easy_ebitengine"
)

func main() {
	k := konstructor.Konstructor{
		EngineWrapper: &easy_ebitengine.EngineWrapper{},
		Domain:        &domain.Domain{},
	}
	k.Init()
}
