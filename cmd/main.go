package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/mmaruf23/golang-rest-api/config"
	"github.com/mmaruf23/golang-rest-api/internal/app/handler"
	"github.com/mmaruf23/golang-rest-api/internal/app/repository"
	"github.com/mmaruf23/golang-rest-api/internal/app/service"
	"github.com/mmaruf23/golang-rest-api/pkg/database"
)

func main() {
	appConfig := config.LoadConfig()

	db := database.NewMySQLConnection(appConfig)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}(db)

	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository(db)

	categoryService := service.NewCategoryService(categoryRepository, validate)

	categoryHanlder := handler.NewCategoryHandler(categoryService)

	router := httprouter.New()

	router.POST("/api/categories", categoryHanlder.Create)
	router.PUT("/api/categories/:categoryId", categoryHanlder.Update)
	router.DELETE("/api/categories/:categoryId", categoryHanlder.Delete)
	router.GET("/api/categories/:categoryId", categoryHanlder.FindById)
	router.GET("/api/categories", categoryHanlder.FindAll)

	port := config.LoadConfig().ServerPort
	fmt.Printf("Server running on PORT %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
