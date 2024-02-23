package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/arraysArrais/api-products/internal/database"
	"github.com/arraysArrais/api-products/internal/service"
	"github.com/arraysArrais/api-products/internal/webserver"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go-api-products")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	port := ":8080"

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCateGoryService(*categoryDB)

	productDB := database.NewProductDB(db)
	productService := service.NewProductService(*productDB)

	WebCategoryHandler := webserver.NewWebCategoryHandler(categoryService)
	WebProductHandler := webserver.NewWebProductHandler(productService)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Get("/category/{id}", WebCategoryHandler.GetCategory)
	router.Get("/category", WebCategoryHandler.GetCategories)
	router.Post("/category", WebCategoryHandler.CreateCategory)

	router.Get("/product/{id}", WebProductHandler.GetProduct)
	router.Get("/product", WebProductHandler.GetProducts)
	router.Get("/product/category/{categoryID}", WebProductHandler.GetProductByCategoryID)
	router.Post("/prodct", WebProductHandler.CreateProduct)

	fmt.Println("Server is running on port ", port)

	http.ListenAndServe(port, router)
}
