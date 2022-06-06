package service

import (
	"errors"
	"no-zero-day/CONTINUE/chapter7_5/entity"
	"no-zero-day/CONTINUE/chapter7_5/repository"
)

type CategoryService struct {
	// name type
	Repository repository.CategoryRepository
}

// create a method extends from CategoryService struct
func (service CategoryService) Get(id string) (*entity.Category, error) {
	// service. get from the method extends
	category := service.Repository.FindById(id)

	if category != nil {
		return category, nil
	} else {
		return nil, errors.New("category not found")
	}
}
