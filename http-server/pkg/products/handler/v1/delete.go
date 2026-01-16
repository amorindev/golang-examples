package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/pkg/products/repository/memory"
	"example.com/shared/api/core"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.ErrorMsg{Msg: "missing product id"})
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.ErrorMsg{Msg: "Invalid id"})
		return
	}

	err = memory.Delete(productId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.ErrorMsg{Msg: err.Error()})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}