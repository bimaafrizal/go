package main

import (
	"belajar-api/app"
	"belajar-api/controller"
	"belajar-api/helper"
	"belajar-api/middleware"
	"belajar-api/repository"
	"belajar-api/service"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)
	auth := middleware.NewAuthMiddleware(router)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: auth,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
