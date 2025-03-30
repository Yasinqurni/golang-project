package repository

import "golang-project/pkg/config"

type Repository struct {
	ExampleRepository
}

func NewRepository(cfg *config.Config) *Repository {
	var (
		exampleRepository = NewExampleRepository(cfg.GormDB)
	)

	return &Repository{
		ExampleRepository: exampleRepository,
	}
}
