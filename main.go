package main

import (
	"belajar_golang_rest_full_api/app"
	"belajar_golang_rest_full_api/controller"
	"belajar_golang_rest_full_api/helper"
	"belajar_golang_rest_full_api/repository"
	"belajar_golang_rest_full_api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()
	router.GET("/api/categories", categoryController.FindAll)
	router.POST("/api/categories", categoryController.Create)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)
	router.PUT("/api/categories/:categoryId", categoryController.Update)

	server := http.Server{
		Addr:    "localhost:8083",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
