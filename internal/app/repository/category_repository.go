package repository

import "github.com/mmaruf23/golang-rest-api/internal/domain"

type CategoryRepository interface {
	Save(category domain.Category) domain.Category
	Update(category domain.Category) domain.Category
	Delete(categoryId int)
	FindById(categoryId int) (domain.Category, error)
	FindAll() []domain.Category
}
