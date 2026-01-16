package v1

import (
	"encoding/json"
	"net/http"

	"example.com/pkg/products/repository/memory"
	"example.com/shared/api/core"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categoryValue := r.FormValue("category")

	products, err := memory.GetAll(categoryValue)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.ErrorMsg{Msg: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
