package handler

import (
	"encoding/json"
	"net/http"

	"example.com/pkg/products/core"
	"example.com/pkg/products/domain"
	"example.com/pkg/products/repository/memory"
	sharedC "example.com/shared/api/core"
)

// CreateProduct godoc
// @Summary      Create a product
// @Description  Create a product with the input payload
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        productCreateReq  body  core.ProductCreateReq  true  "Create product"
// @Success      201  {object}  domain.Product
// @Failure      400  {object}  core.ErrorMsg
// @Failure      500  {object}  core.ErrorMsg
// @Router       /products [post]
func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req core.ProductCreateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(sharedC.ErrorMsg{Msg: "Invalid request body"})
		return
	}

	defer r.Body.Close()

	if req.CategoryName == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(sharedC.ErrorMsg{
			Msg: "category_name is required",
		})
		return
	}

	product := &domain.Product{
		Name:         req.Name,
		Desc:         req.Desc,
		Price:        req.Price,
		Stock:        req.Stock,
		CategoryName: req.CategoryName,
	}

	err := memory.Create(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(sharedC.ErrorMsg{Msg: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(req)
}
