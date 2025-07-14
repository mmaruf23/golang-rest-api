package service

import "github.com/mmaruf23/golang-rest-api/internal/domain"

type CategoryService interface {
	Create(category domain.Category) domain.Category
	Update(category domain.Category) domain.Category
	Delete(categoryId int)
	FindById(categoryId int) (domain.Category, error)
	FindAll() []domain.Category
}
