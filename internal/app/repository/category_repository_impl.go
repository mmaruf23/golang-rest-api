package repository

import (
	"database/sql"
	"fmt"

	"github.com/mmaruf23/golang-rest-api/internal/domain"
	"github.com/mmaruf23/golang-rest-api/internal/helper"
)

type CategoryRepositoryImpl struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &CategoryRepositoryImpl{DB: db}
}

func (repository *CategoryRepositoryImpl) Save(category domain.Category) domain.Category {
	query := "INSERT INTO categories(name) VALUES (?)"
	stmt, err := repository.DB.Prepare(query)

	helper.PanicIfError(err)
	defer stmt.Close()

	// kalau mau pake timestamp update bisa disini sekarang mah kagak dulu

	result, err := stmt.Exec(category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	fmt.Printf("Category saved with ID: %d\n", category.Id)

	return category
}

func (repository *CategoryRepositoryImpl) Update(category domain.Category) domain.Category {
	query := "UPDATE categories SET name = ? WHERE id = ?"
	stmt, err := repository.DB.Prepare(query)
	helper.PanicIfError(err)
	defer stmt.Close()

	_, err = stmt.Exec(category.Name, category.Id)
	helper.PanicIfError(err)

	fmt.Printf("Category Updated with ID: %d\n", category.Id)
	return category
}

func (repository *CategoryRepositoryImpl) Delete(categoryId int) {
	query := "DELETE from categories WHERE id = ?"
	stmt, err := repository.DB.Prepare(query)
	helper.PanicIfError(err)
	defer stmt.Close()

	_, err = stmt.Exec(categoryId)
	helper.PanicIfError(err)

	fmt.Printf("Category Deleted with ID: %d\n", categoryId)
}

func (repository *CategoryRepositoryImpl) FindById(categoryId int) (domain.Category, error) {
	query := "SELECT id, name FROM categories WHERE id = ?"
	rows, err := repository.DB.Query(query, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, fmt.Errorf("product with ID %d not found", categoryId)
	}
}

func (repository *CategoryRepositoryImpl) FindAll() []domain.Category {
	query := "SELECT id, name FROM categories"
	rows, err := repository.DB.Query(query)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}
