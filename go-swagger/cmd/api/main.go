package main

import (
	"log"
	"net/http"

	_ "example.com/cmd/api/api-docs"
	"example.com/pkg/products/handler"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title        My API with net/http
// @version      1.0
// @description  Example project demonstrating Swagger with net/http.
// @host         localhost:8090
// @BasePath     /
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /products", handler.GetAll)
	mux.HandleFunc("GET /products/{id}", handler.Get)
	mux.HandleFunc("POST /products", handler.Create)
	mux.HandleFunc("PUT /products/{id}", handler.Update)
	mux.HandleFunc("DELETE /products/{id}", handler.Delete)
	mux.HandleFunc("POST /products/{productId}/img", handler.UploadImage)

	mux.HandleFunc("/docs/", httpSwagger.WrapHandler)

	if err := http.ListenAndServe(":8090", mux); err != nil {
		log.Fatal(err)
	}
}
