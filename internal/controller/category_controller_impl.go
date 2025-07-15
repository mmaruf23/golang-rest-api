package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/mmaruf23/golang-rest-api/internal/helper"
	"github.com/mmaruf23/golang-rest-api/internal/model/web"
	"github.com/mmaruf23/golang-rest-api/internal/service"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{CategoryService: categoryService}
}

func (controller *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}

	helper.ReadFromRequestBody(r, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(r.Context(), categoryCreateRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (controller *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(r, &categoryUpdateRequest)

	categoryId := p.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(r.Context(), categoryUpdateRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryId := p.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	if err != nil {
		helper.PanicIfError(err)
	}

	controller.CategoryService.Delete(r.Context(), id)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   "Succesfully Deleted",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryId := p.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	if err != nil {
		helper.PanicIfError(err)
	}

	categoryResponse := controller.CategoryService.FindById(r.Context(), id)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryResponses := controller.CategoryService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}
