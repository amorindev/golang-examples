package products

import (
	"api-rest-crud/internal/entities"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categoryValue := r.FormValue("category")

	products, err := getProductsDto(categoryValue)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entities.ErrorMessage{Msg: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")
	productId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entities.ErrorMessage{Msg: "Invalid id"})
		return
	}

	product, err := getProductDto(productId)
	if err != nil {
		if (err == productNotFound{}) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(entities.ErrorMessage{Msg: err.Error()})
		return
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}


func PostProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entities.ErrorMessage{Msg: "Failed to read request body"})
		return
	}

	var product Product
	err = json.Unmarshal(body, &product)
	if err != nil {
	  	w.WriteHeader(http.StatusBadRequest)
	 	json.NewEncoder(w).Encode(entities.ErrorMessage{Msg: "Failed to parse request body"})
		return
	}
	
	err = createProductDto(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entities.ErrorMessage{Msg: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nil)
}

func PutProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	id := r.PathValue("id")
	productId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entities.ErrorMessage{Msg: "Invalid id"})
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entities.ErrorMessage{Msg: "Failed to read request body"})
		return
	}

	var newProduct Product
	err = json.Unmarshal(body,&newProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entities.ErrorMessage{Msg: "Failed to parse body"})
		return
	}

	err = updateProductDto(productId,newProduct)
	if err != nil {
		if (err == productNotFound{}) {
			w.WriteHeader(http.StatusNotFound)
		} else{
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(entities.ErrorMessage{Msg: err.Error()})
		return
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(nil)
} 

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")

	productId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entities.ErrorMessage{Msg: "Invalid id"})
		return
	}

	err = deleteProductDto(productId)
	if err != nil {
		if (err == productNotFound{}) {
			w.WriteHeader(http.StatusNotFound)
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entities.ErrorMessage{Msg: err.Error()})
		return
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(nil)
}


