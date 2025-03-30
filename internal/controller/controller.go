package controller

import "golang-project/internal/service"

type Controller struct {
	ExampleController
}

func NewController(s *service.Service) *Controller {
	var (
		exampleController = NewExampleController(s.ExampleService)
	)

	return &Controller{
		ExampleController: exampleController,
	}
}
