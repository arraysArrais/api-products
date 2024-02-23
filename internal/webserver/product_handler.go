package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/arraysArrais/api-products/internal/entity"
	"github.com/arraysArrais/api-products/internal/service"
	"github.com/go-chi/chi/v5"
)

type WebProductHandler struct {
	ProductService *service.ProductService
}

func NewWebProductHandler(productService *service.ProductService) *WebProductHandler {
	return &WebProductHandler{ProductService: productService}
}

func (wph *WebProductHandler) GetProducts(writer http.ResponseWriter, request *http.Request) {
	products, err := wph.ProductService.GetProducts()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(products)
}

func (wph *WebProductHandler) GetProduct(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")

	if id == "" {
		http.Error(writer, "id is required", http.StatusBadRequest)
		return
	}
	product, err := wph.ProductService.GetProduct(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(product)
}

func (wph *WebProductHandler) GetProductByCategoryID(writer http.ResponseWriter, request *http.Request) {
	categoryID := chi.URLParam(request, "categoryID")
	if categoryID == "" {
		http.Error(writer, "id is required", http.StatusBadRequest)
		return
	}
	products, err := wph.ProductService.GetProductByCategoryID(categoryID)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(products)
}

func (wph *WebProductHandler) CreateProduct(writer http.ResponseWriter, request *http.Request) {
	var product entity.Product

	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := wph.ProductService.CreateProduct(product.Name, product.Description, product.CategoryID, product.ImageURL, product.Price)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(result)
}
