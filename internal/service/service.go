package service

import "golang-project/internal/repository"

type Service struct {
	ExampleService
}

func NewService(r *repository.Repository) *Service {
	var (
		exampleService = NewExampleService(r.ExampleRepository)
	)
	return &Service{
		ExampleService: exampleService,
	}
}
