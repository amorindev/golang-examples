package memory

import "example.com/pkg/products/domain"

func Create(product *domain.Product) error {
	if len(products) == 0 {
		product.Id = 1
	} else {
		pLength := len(products)
		product.Id = products[pLength-1].Id + 1
	}

	products = append(products, product)
	return nil
}
