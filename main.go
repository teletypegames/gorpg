package main

import (
	"git.teletypegames.org/engines/gorpg/src/domain"
	"git.teletypegames.org/engines/gorpg/src/konstructor"
	"git.teletypegames.org/engines/gorpg/src/konstructor/engines/easy_ebitengine"
)

func main() {
	k := konstructor.Konstructor{
		EngineWrapper: &easy_ebitengine.EngineWrapper{},
		Domain:        &domain.Domain{},
	}
	k.Init()
}
