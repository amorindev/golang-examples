package main

import (
	"log"
	"net/http"

	mdw "example.com/shared/api/middlewares"

	_ "example.com/cmd/api/api-docs"

	productV1Handler "example.com/pkg/products/handler/v1"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title        My API with net/http
// @version      1.0
// @description  Example project demonstrating Swagger with net/http.
// @host         localhost:8082
// @BasePath     /api/v1
func main() {
	mux := http.NewServeMux()

	// Apply CORS then logging middleware to mux
	//muxWithMdw := mdw.MyHandlerMiddleware(mux)
	muxWithMdw := mdw.LogRequest(mdw.CorsMiddleware(mux))

	v1 := http.NewServeMux()

	v1.HandleFunc("GET /products", productV1Handler.GetAll)
	v1.HandleFunc("GET /products/{id}", productV1Handler.Get)
	v1.HandleFunc("POST /products", productV1Handler.Create)
	v1.HandleFunc("PUT /products/{id}", productV1Handler.Update)
	v1.HandleFunc("DELETE /products/{id}", productV1Handler.Delete)

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", v1))

	//v2 := http.NewServeMux()

	//v2.HandleFunc("GET /products", productV2Handler.GetAll)
	//v2.HandleFunc("GET /products/{id}", productV2Handler.Get)
	//v2.HandleFunc("POST /products", productV2Handler.Create)
	//v2.HandleFunc("PUT /products/{id}", productV2Handler.Update)
	//v2.HandleFunc("DELETE /products/{id}", productV2Handler.Delete)

	//mux.Handle("/api/v2/", http.StripPrefix("/api/v2", v2))

	// Health check route
	//mux.HandleFunc("GET /ping", mdw.MyMiddleware(handler.Ping))
	//mux.HandleFunc("GET /ping", sharedH.Ping)

	// Api docs
	mux.HandleFunc("/docs/", httpSwagger.WrapHandler)

	// Templates
	//mux.HandleFunc("GET /admin/home", adminH.HomePage)

	if err := http.ListenAndServe(":8082", muxWithMdw); err != nil {
		log.Fatal(err)
	}
}
