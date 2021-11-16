package controller

import (
	"net/http"
	"strconv"

	"github.com/khilmi-aminudin/golang_restful_api/helper"
	"github.com/khilmi-aminudin/golang_restful_api/model/web"
	"github.com/khilmi-aminudin/golang_restful_api/service"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type categoryController struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &categoryController{
		CategoryService: categoryService,
	}
}

func (controller *categoryController) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var categoryCreateRequest web.CategoryCreateRequest
	helper.ReadFromRequestBody(r, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(r.Context(), categoryCreateRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteResponseBody(w, webResponse)
}

func (controller *categoryController) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var categoryUpdateRequest web.CategoryUpdateRequest
	helper.ReadFromRequestBody(r, &categoryUpdateRequest)

	categoryId := params.ByName("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(r.Context(), categoryUpdateRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteResponseBody(w, webResponse)
}

func (controller *categoryController) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helper.WriteResponseBody(w, webResponse)
}

func (controller *categoryController) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteResponseBody(w, webResponse)
}

func (controller *categoryController) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	categoryResponses := controller.CategoryService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.WriteResponseBody(w, webResponse)
}
