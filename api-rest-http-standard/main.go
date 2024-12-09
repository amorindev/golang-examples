package main

import (
	"api-rest-crud/internal/services/products"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	v1 := http.NewServeMux()

	v1.HandleFunc("GET /products", products.GetProducts)
	v1.HandleFunc("GET /products/{id}", products.GetProduct)
	v1.HandleFunc("POST /products", products.PostProduct)
	v1.HandleFunc("PUT /products/{id}", products.PutProduct)
	v1.HandleFunc("DELETE /products/{id}", products.DeleteProduct)

	mux.Handle("/v1/", http.StripPrefix("/v1", v1))

	v2 := http.NewServeMux()
	v2.HandleFunc("GET /products", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Get products v2"}`))
	})

	v2.HandleFunc("DELETE /products/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "appication/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Product deleted v2"}`))
	})

	mux.Handle("/v2/", http.StripPrefix("/v2",v2))

	if err := http.ListenAndServe(":8082", mux); err != nil {
		log.Fatal(err)
	}
}


// apra mantener un sistema de versiones para la api usamos .............. 
// prefijo v1
// antes de que se envíe al proximo enrutador
// a	

