package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/mmaruf23/golang-rest-api/internal/app/service"
	"github.com/mmaruf23/golang-rest-api/internal/domain"
	"github.com/mmaruf23/golang-rest-api/internal/helper"
)

type CategoryHandler interface {
	Create(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}

type CategoryHandlerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) CategoryHandler {
	return &CategoryHandlerImpl{CategoryService: categoryService}
}

func (handler *CategoryHandlerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryRequest := domain.Category{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&categoryRequest)

	// nanti ubah , jangan panic di handler - katanya.
	helper.PanicIfError(err)

	newCategory := handler.CategoryService.Create(categoryRequest)

	webResponse := domain.WebResponse{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   newCategory,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)

}

// todo : tambahin handler lain
// update, delete, findbyid, findall

func (handler *CategoryHandlerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryRequest := domain.Category{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&categoryRequest)

	// nanti ubah , jangan panic di handler - katanya.
	helper.PanicIfError(err)

	categoryId := p.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	if err != nil {
		helper.PanicIfError(fmt.Errorf("invalid category ID : %v", err))
	}

	categoryRequest.Id = id

	updatedCategory := handler.CategoryService.Update(categoryRequest)

	webReponse := domain.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   updatedCategory,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webReponse)
	helper.PanicIfError(err)
}

func (handler *CategoryHandlerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryId := p.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	if err != nil {
		helper.PanicIfError(fmt.Errorf("invalid category ID : %v", err))
	}

	handler.CategoryService.Delete(id)

	webResponse := domain.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   "Succesfully Deleted",
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (handler *CategoryHandlerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryId := p.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	if err != nil {
		helper.PanicIfError(fmt.Errorf("invalid category ID : %v", err))
	}

	categoryResponse, err := handler.CategoryService.FindById(id)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		encoder := json.NewEncoder(w)

		errorReponse := domain.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   err.Error(),
		}
		encoder.Encode(errorReponse)
		return
	}

	webResponse := domain.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)

}

func (handler *CategoryHandlerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categories := handler.CategoryService.FindAll()
	webResponse := domain.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categories,
	}
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
