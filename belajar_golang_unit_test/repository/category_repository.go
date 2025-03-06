package repository

import "belajar_golang_unit_test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
	// FindAll() ([]*entity.Category, error)
	// Save(category *entity.Category) error
	// Update(category *entity.Category) error
	// Delete(category *entity.Category) error
}
