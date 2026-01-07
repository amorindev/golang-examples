package memory

import (
	"example.com/pkg/products/domain"
	"errors"
)

func Update(id int, newProduct *domain.Product) error {
	for i, product := range products {
		if product.Id == id {
			newProduct.Id = id
			products[i] = newProduct
			return nil
		}
	}
    return errors.New("product not found")
}