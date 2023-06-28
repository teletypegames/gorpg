package main

import (
	"game/domain"
	"game/presenter"
)

func main() {
	context := domain.CreateContext()

	domain := domain.Domain{
		Context: context,
		Presenter: presenter.Presenter{
			Context: context,
		},
	}
	domain.Init()
}
