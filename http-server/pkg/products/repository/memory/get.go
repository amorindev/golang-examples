package memory

import (
	"example.com/pkg/products/domain"
	"errors"
)

func Get(id int) (*domain.Product, error) {
	for _, product := range products {
		if product.ID == id {
			return product, nil
		}
	}
	return nil, errors.New("product not found")
}
