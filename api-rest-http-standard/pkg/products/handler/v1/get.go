package handler

import (
	"example.com/pkg/products/repository/memory"
	"example.com/shared/api/core"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// Get godoc
// @Summary      Get product by ID
// @Tags         products
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  domain.Product
// @Failure      400  {object}  core.ErrorMsg
// @Router       /products/{id} [get]
func Get(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)

    w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")
	productId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.ErrorMsg{Msg: "Invalid id"})
		return
	}

	product, err := memory.Get(productId)
	if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(core.ErrorMsg{Msg: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}