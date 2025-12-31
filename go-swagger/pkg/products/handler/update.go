package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/pkg/products/domain"
	"example.com/pkg/products/repository/memory"
	"example.com/shared/api/core"
)

// Update godoc
// @Summary      Update product
// @Description  Update an existing product by its ID using the provided JSON body
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id       path      string             true  "Product ID"
// @Param        product  body      domain.Product  true  "Updated product"
// @Success      200  {object}  domain.Product
// @Failure      400  {object}  core.ErrorMsg
// @Failure      500  {object}  core.ErrorMsg
// @Router       /products/{id} [put]
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
