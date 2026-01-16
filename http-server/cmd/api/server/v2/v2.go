package v2

import (
	"net/http"

	v2Handler "example.com/pkg/products/handler/v2"
	"example.com/shared/api/handler"
)

func New() http.Handler {
	mux := http.NewServeMux()

	// Api version
	v2 := http.NewServeMux()
	mux.Handle("/v2/", http.StripPrefix("/v2", v2))

	// Here's everything your app needs
	// database connection

	v2.HandleFunc("GET /products", v2Handler.GetAll)
	v2.HandleFunc("GET /products/{id}", v2Handler.Get)
	v2.HandleFunc("POST /products", v2Handler.Create)
	v2.HandleFunc("PUT /products/{id}", v2Handler.Update)
	v2.HandleFunc("DELETE /products/{id}", v2Handler.Delete)

	mux.HandleFunc("GET /ping", handler.Ping)

	return mux
}
