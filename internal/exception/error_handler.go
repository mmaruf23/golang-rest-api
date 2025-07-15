package exception

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mmaruf23/golang-rest-api/internal/helper"
	"github.com/mmaruf23/golang-rest-api/internal/model/web"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if validationErrors(w, r, err) {
		return
	}

	if notFoundError(w, err) {
		return
	}

	internalServerError(w, r, err)
}

func validationErrors(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		errorResponse := web.ErrorResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD REQUEST",
			Message: exception.Error(),
		}

		helper.WriteToResponseBody(w, errorResponse)
	}

	return ok
}

func notFoundError(w http.ResponseWriter, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		errorResponse := web.ErrorResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD REQUEST",
			Message: exception.Error,
		}

		helper.WriteToResponseBody(w, errorResponse)
	}

	return ok
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	errorResponse := web.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Status:  "INTERNAL SERVER ERROR",
		Message: err,
	}

	helper.WriteToResponseBody(w, errorResponse)
}
