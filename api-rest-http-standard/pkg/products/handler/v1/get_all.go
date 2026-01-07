package handler

import (
	"example.com/pkg/products/repository/memory"
	"example.com/shared/api/core"
	"encoding/json"
	"net/http"
)

// GetAll godoc
// @Summary      List products
// @Tags         products
// @Produce      json
// @Param        category  header    string  false  "Filter by category"
// @Success      200  {array}   domain.Product
// @Failure      500  {object}  core.ErrorMsg
// @Router       /products [get]
func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categoryHeader := r.Header.Get("category")

	products, err := memory.GetAll(categoryHeader)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(core.ErrorMsg{Msg: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
