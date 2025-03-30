package service

import (
	"golang-project/internal/entity/model"
	"golang-project/internal/entity/request"
	"golang-project/internal/repository"
)

type ExampleService interface {
	Create(data request.ExampleBodyRequest) ([]uint32, error)
}

type exampleService struct {
	exampleRepository repository.ExampleRepository
}

func NewExampleService(exampleRepository repository.ExampleRepository) ExampleService {
	return &exampleService{
		exampleRepository: exampleRepository,
	}
}

func (s *exampleService) Create(data request.ExampleBodyRequest) ([]uint32, error) {
	var examples []model.Example
	for _, example := range data.Examples {
		examples = append(examples, model.Example{
			Name:       example.Name,
			Email:      example.Email,
			Age:        example.Age,
			Phone:      example.Phone,
			Address:    example.Address,
			Status:     model.ACTIVE,
			IsVerified: false,
		})
	}

	return s.exampleRepository.Create(examples)
}
