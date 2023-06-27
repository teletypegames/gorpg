package main

import (
	"game/controller"
	"game/domain"
	"game/presenter"
)

func main() {
	context := domain.CreateContext()

	controller := controller.Controller{
		Context: context,
	}
	controller.Init()

	presenter := presenter.Presenter{
		Context: context,
	}
	presenter.Init()

	domain := domain.Domain{
		Context: context,
	}
	domain.Init()

	domain.Start()
	controller.Start()
	presenter.Start()
}
