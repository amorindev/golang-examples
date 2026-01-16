package memory

import (
	"example.com/pkg/products/domain"
	"errors"
)

func Update(id int, newProduct *domain.Product) error {
	for i, product := range products {
		if product.ID == id {
			newProduct.ID = id
			products[i] = newProduct
			return nil
		}
	}
    return errors.New("product not found")
}

