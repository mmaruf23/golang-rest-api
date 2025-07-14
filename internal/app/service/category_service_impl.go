package service

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/mmaruf23/golang-rest-api/internal/app/repository"
	"github.com/mmaruf23/golang-rest-api/internal/domain"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		Validate:           validate,
	}
}

// implementasi

func (service *CategoryServiceImpl) Create(category domain.Category) domain.Category {
	err := service.Validate.Struct(category)
	if err != nil {
		// sementara panic dulu
		panic(fmt.Sprintf("Validation error: %v", err))
	}

	newCategory := service.CategoryRepository.Save(category)
	return newCategory
}

func (service *CategoryServiceImpl) Update(category domain.Category) domain.Category {

	if category.Id == 0 {
		panic("Category ID required bambang!")
	}

	_, err := service.CategoryRepository.FindById(category.Id)
	if err != nil {
		panic(fmt.Sprintf("Category with ID : %d not found, cannot update : %v", category.Id, err))
	}

	return service.CategoryRepository.Update(category)

}

func (service *CategoryServiceImpl) Delete(categoryId int) {
	_, err := service.CategoryRepository.FindById(categoryId)
	if err != nil {
		panic(fmt.Sprintf("Category with ID %d not found for deletion: %v", categoryId, err))
	}

	service.CategoryRepository.Delete(categoryId)
}

func (service *CategoryServiceImpl) FindById(categoryId int) (domain.Category, error) {
	category, err := service.CategoryRepository.FindById(categoryId)
	if err != nil {
		return category, fmt.Errorf("could not find product by ID: %w", err)
	}
	return category, nil
}
func (service *CategoryServiceImpl) FindAll() []domain.Category {
	return service.CategoryRepository.FindAll()
}
