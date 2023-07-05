package entity

type EngineOptions struct {
	Domain   DomainInterface
	KContext *KContext
	Settings *Settings
}

type EngineWrapperInterface interface {
	Init(options EngineOptions)
	Run()
}
