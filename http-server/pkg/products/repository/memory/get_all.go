package memory

import (
	"example.com/pkg/products/domain"
)

func GetAll(category string) ([]*domain.Product, error) {
	if category == "" {
		return products, nil
	}
	var filteredProducts []*domain.Product
	for _, product := range products {
		if product.CategoryName == category {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return products, nil
}
