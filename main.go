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
	presenter := presenter.Presenter{
		Context: context,
	}
	domain := domain.Domain{
		Context:    context,
		Controller: controller,
		Presenter:  presenter,
	}
	domain.Init()
}
