package konstructor

import (
	"game/konstructor/entity"
)

type Konstructor struct {
	Domain        entity.DomainInterface
	KContext      *entity.KContext
	Settings      *entity.Settings
	EngineWrapper entity.EngineWrapperInterface
}

func (k Konstructor) Init() {
	k.Domain.Init()
	k.EngineWrapper.Init(entity.EngineOptions{
		Domain:   k.Domain,
		KContext: k.KContext,
		Settings: k.Settings,
	})
	k.EngineWrapper.Run()
}
