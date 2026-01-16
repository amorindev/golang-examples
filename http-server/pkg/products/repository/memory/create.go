package memory

import "example.com/pkg/products/domain"

func Create(product *domain.Product) error {
	if len(products) == 0 {
		product.ID = 1
	} else {
		pLength := len(products)
		product.ID = products[pLength-1].ID + 1
	}

	products = append(products, product)
	return nil
}