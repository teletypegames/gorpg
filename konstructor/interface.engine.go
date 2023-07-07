package konstructor

type EngineArgs struct {
	Domain   DomainInterface
	KContext *KContext
	Config   *Config
}

type EngineWrapperInterface interface {
	Init(options EngineArgs)
	Run()
}
