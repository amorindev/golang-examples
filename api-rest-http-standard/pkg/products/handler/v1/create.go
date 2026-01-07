package handler

import (
	"example.com/pkg/products/domain"
	"example.com/pkg/products/repository/memory"
	"example.com/shared/api/core"
	"encoding/json"
	"net/http"
)

// Create godoc
// @Summary      Create product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body  domain.Product  true  "Product data"
// @Success      201  {object}  domain.Product
// @Failure      400  {object}  core.ErrorMsg
// @Failure      500  {object}  core.ErrorMsg
// @Router       /products [post]
func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req *domain.Product
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(core.ErrorMsg{Msg: "Invalid request body"})
		return
	}

	defer r.Body.Close()

	// Validations

	err := memory.Create(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(core.ErrorMsg{Msg: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(req)
}
