package entity

type EngineArgs struct {
	Domain   DomainInterface
	KContext *KContext
	Settings *Settings
}

type EngineWrapperInterface interface {
	Init(options EngineArgs)
	Run()
}
