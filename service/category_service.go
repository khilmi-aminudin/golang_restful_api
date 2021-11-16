package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/khilmi-aminudin/golang_restful_api/helper"
	"github.com/khilmi-aminudin/golang_restful_api/model/domain"
	"github.com/khilmi-aminudin/golang_restful_api/model/web"
	"github.com/khilmi-aminudin/golang_restful_api/repository"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}

type categoryService struct {
	Repository repository.CategoryRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewCategoryService(repository repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return &categoryService{
		Repository: repository,
		DB:         db,
		Validate:   validate,
	}
}

func (service *categoryService) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.Repository.Save(ctx, tx, category)

	response := web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
	return response
}

func (service *categoryService) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.Repository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	category.Name = request.Name

	category = service.Repository.Update(ctx, tx, category)

	response := web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
	return response
}

func (service *categoryService) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.Repository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	service.Repository.Delete(ctx, tx, category)

}

func (service *categoryService) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.Repository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	response := web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}

	return response
}

func (service *categoryService) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.Repository.FindAll(ctx, tx)

	var responses []web.CategoryResponse

	for _, category := range categories {
		categoryResponse := web.CategoryResponse{
			Id:   category.Id,
			Name: category.Name,
		}
		responses = append(responses, categoryResponse)
	}

	return responses
}
