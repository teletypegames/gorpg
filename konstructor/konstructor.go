package konstructor

type Konstructor struct {
	Domain        DomainInterface
	EngineWrapper EngineWrapperInterface
}

func (k Konstructor) Init() {
	k.Domain.Init()
	k.EngineWrapper.Init(EngineArgs{
		Domain:   k.Domain,
		KContext: k.Domain.GetInitialKContext(),
		Config:   k.Domain.GetInitialConfig(),
	})
	k.EngineWrapper.Run()
}
