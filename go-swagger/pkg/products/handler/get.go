package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/pkg/products/repository/memory"
	"example.com/shared/api/core"
)

// Get godoc
// @Summary      Get a product
// @Description get a product with id param
// @Tags         products
// @Accept      json
// @Produce      json
// @Param        id   path      string  true  "Product ID"
// @Success      200  {object}  domain.Product
// @Failure      400  {object}  core.ErrorMsg
// @Failure      500  {object}  core.ErrorMsg
// @Router       /products/{id} [get]
func Get(w http.ResponseWriter, r *http.Request) {
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
