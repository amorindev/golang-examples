package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/pkg/products/domain"
	"example.com/pkg/products/repository/memory"
	"example.com/shared/api/core"
)

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")
	productId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.ErrorMsg{Msg: "Invalid id"})
		return
	}

	var req *domain.Product
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.ErrorMsg{Msg: "Invalid request body"})
		return
	}
	defer r.Body.Close()

	// validations

	err = memory.Update(productId, req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.ErrorMsg{Msg: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(req)
}