package main

import (
	"fmt"
	"golang-restapi-noframework/config"
	"golang-restapi-noframework/controller"
	"golang-restapi-noframework/helper"
	"golang-restapi-noframework/repository"
	"golang-restapi-noframework/router"
	"golang-restapi-noframework/service"
	"net/http"
)

func main() {
	fmt.Printf("Start server")
	// database
	db := config.DatabaseConnection()

	// repository
	productRepository := repository.NewProductRepository(db)

	// service
	productService := service.NewProductServiceImpl(productRepository)

	// controller
	productController := controller.NewProductController(productService)

	// router
	routes := router.NewRouter(productController)

	server := http.Server{Addr: "localhost:8888", Handler: routes}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
