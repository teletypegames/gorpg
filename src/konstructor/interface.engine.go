package konstructor

// Initial arguments of an engine
type EngineArgs struct {
	Domain   DomainInterface
	KContext *KContext
	Config   *Config
}

// Engine wrapper interface. You need to implement to this for your EntityWrapper, which can run your game engine
type EngineWrapperInterface interface {
	Init(options EngineArgs)
	Run()
}
