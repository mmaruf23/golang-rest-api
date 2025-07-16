package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mmaruf23/golang-rest-api/config"
	"github.com/mmaruf23/golang-rest-api/internal/app"
	"github.com/mmaruf23/golang-rest-api/internal/controller"
	"github.com/mmaruf23/golang-rest-api/internal/helper"
	"github.com/mmaruf23/golang-rest-api/internal/middleware"
	"github.com/mmaruf23/golang-rest-api/internal/repository"
	"github.com/mmaruf23/golang-rest-api/internal/service"
)

func main() {
	appConfig := config.LoadConfig()

	db := app.NewMySQLConnection(appConfig)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}(db)

	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	port := config.LoadConfig().ServerPort

	server := http.Server{
		Addr:    "localhost:" + port,
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
