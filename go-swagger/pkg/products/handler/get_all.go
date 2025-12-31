package handler

import (
	"encoding/json"
	"net/http"

	"example.com/pkg/products/repository/memory"
	"example.com/shared/api/core"
)

// GetAll godoc
// @Summary      List products
// @Description  Returns a list of products. If the optional 'category' header is provided, the result will be filtered by that category.
// @Tags         products
// @Produce      json
// @Param        category  header    string  false  "Filter by category"
// @Success      200  {array}   domain.Product
// @Failure      400  {object}  core.ErrorMsg
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
