package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"github.com/mmaruf23/golang-rest-api/config"
	"github.com/mmaruf23/golang-rest-api/internal/app"
	"github.com/mmaruf23/golang-rest-api/internal/controller"
	"github.com/mmaruf23/golang-rest-api/internal/middleware"
	"github.com/mmaruf23/golang-rest-api/internal/repository"
	"github.com/mmaruf23/golang-rest-api/internal/service"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepositoryImpl,
	wire.Bind(new(repository.CategoryRepository), (new(*repository.CategoryRepositoryImpl))),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), (new(*service.CategoryServiceImpl))),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), (new(*controller.CategoryControllerImpl))),
)

func InitializeServer(appConfig *config.AppConfig) *http.Server {
	wire.Build(
		app.NewMySQLConnection,
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), (new(*httprouter.Router))),
		middleware.NewAuthMiddleware,
		NewServer)
	return nil
}
