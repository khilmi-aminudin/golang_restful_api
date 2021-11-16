package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/khilmi-aminudin/golang_restful_api/app"
	"github.com/khilmi-aminudin/golang_restful_api/controller"
	"github.com/khilmi-aminudin/golang_restful_api/helper"
	"github.com/khilmi-aminudin/golang_restful_api/repository"
	"github.com/khilmi-aminudin/golang_restful_api/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db                 *sql.DB                       = app.NewDB()
	validate           *validator.Validate           = validator.New()
	categoryRepository repository.CategoryRepository = repository.NewCategoryRepository()
	categoryService    service.CategoryService       = service.NewCategoryService(categoryRepository, db, validate)
	categoryController controller.CategoryController = controller.NewCategoryController(categoryService)
)

func main() {

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:id", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:id", categoryController.Update)
	router.DELETE("/api/categories/:id", categoryController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	fmt.Println("app is running on http://localhost:3000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
