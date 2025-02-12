package repository

import "golang_software_testing/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}