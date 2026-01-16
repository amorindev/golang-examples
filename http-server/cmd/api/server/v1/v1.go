package v1

import (
	"net/http"

	"example.com/shared/api/handler"
	v1Handler "example.com/pkg/products/handler/v1"
)

func New() http.Handler {
	mux := http.NewServeMux()

	// Api version
	v1 := http.NewServeMux()
	mux.Handle("/v1/", http.StripPrefix("/v1", v1))

	// Here's everything your app needs
	// database connection

	v1.HandleFunc("GET /products", v1Handler.GetAll)
	v1.HandleFunc("GET /products/{id}", v1Handler.Get)
	v1.HandleFunc("POST /products", v1Handler.Create)
	v1.HandleFunc("PUT /products/{id}", v1Handler.Update)
	v1.HandleFunc("DELETE /products/{id}", v1Handler.Delete)

	mux.HandleFunc("/ping", handler.Ping)

	return mux
}
