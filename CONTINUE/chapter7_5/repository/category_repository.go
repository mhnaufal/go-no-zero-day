package repository

import "no-zero-day/CONTINUE/chapter7_5/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
