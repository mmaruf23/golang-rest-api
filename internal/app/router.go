package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mmaruf23/golang-rest-api/internal/controller"
	"github.com/mmaruf23/golang-rest-api/internal/exception"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.GET("/api/categories", categoryController.FindAll)

	router.PanicHandler = exception.ErrorHandler

	return router
}
