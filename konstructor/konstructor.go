package konstructor

type Konstructor struct {
	Domain        DomainInterface
	KContext      *KContext
	Settings      *Settings
	EngineWrapper EngineWrapperInterface
}

func (k Konstructor) Init() {
	k.Domain.Init()
	k.EngineWrapper.Init(EngineArgs{
		Domain:   k.Domain,
		KContext: k.KContext,
		Settings: k.Settings,
	})
	k.EngineWrapper.Run()
}
