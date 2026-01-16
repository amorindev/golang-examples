package v1

import (
	"encoding/json"
	"net/http"

	"example.com/pkg/products/core"
	"example.com/pkg/products/domain"
	"example.com/pkg/products/repository/memory"
	sharedC "example.com/shared/api/core"
)

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req *core.ProductCreateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(sharedC.ErrorMsg{Msg: "Invalid request body"})
		return
	}

	defer r.Body.Close()

	// Validations here

	product := &domain.Product{
		Name:         req.Name,
		Desc:         req.Desc,
		Price:        req.Price,
		Stock:        req.Stock,
		CategoryName: req.CategoryName,
	}

	err := memory.Create(product)
	if err != nil {
		// manage error
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(sharedC.ErrorMsg{Msg: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(req)
}
