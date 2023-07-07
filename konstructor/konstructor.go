package konstructor

type Konstructor struct {
	Domain        DomainInterface
	KContext      *KContext
	Config        *Config
	EngineWrapper EngineWrapperInterface
}

func (k Konstructor) Init() {
	k.Domain.Init()
	k.EngineWrapper.Init(EngineArgs{
		Domain:   k.Domain,
		KContext: k.KContext,
		Config:   k.Config,
	})
	k.EngineWrapper.Run()
}
