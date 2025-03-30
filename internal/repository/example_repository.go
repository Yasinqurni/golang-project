package repository

import (
	"golang-project/internal/entity/model"

	"gorm.io/gorm"
)

type ExampleRepository interface {
	Create(examples []model.Example) ([]uint32, error)
}

type exampleRepository struct {
	db *gorm.DB
}

func NewExampleRepository(db *gorm.DB) ExampleRepository {
	return &exampleRepository{
		db: db,
	}
}

func (e *exampleRepository) Create(examples []model.Example) ([]uint32, error) {
	result := e.db.Create(&examples)
	if result.Error != nil {
		return nil, result.Error
	}

	var ids []uint32
	for _, example := range examples {
		ids = append(ids, uint32(example.ID))
	}

	return ids, nil
}
