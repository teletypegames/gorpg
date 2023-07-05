package konstructor

import (
	"game/konstructor/entity"
)

type Konstructor struct {
	Domain        entity.DomainInterface
	Settings      entity.Settings
	KContext      *entity.KContext
	EngineWrapper entity.EngineWrapperInterface
}

func (k Konstructor) Init() {
	k.Domain.Init()
	k.EngineWrapper.Init(entity.EngineOptions{
		KContext: *k.KContext,
		Domain:   k.Domain,
		Settings: k.Settings,
	})
	k.EngineWrapper.Run()
}
