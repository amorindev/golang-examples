package memory

import (
	"example.com/pkg/products/domain"
	"strings"
)

func GetAll(category string) ([]*domain.Product, error) {
	if strings.Compare(category, "Technology") == 0 {
		var filteredProducts []*domain.Product
		for _, product := range products {
			if strings.Compare(product.CategoryName, category) == 0 {
				filteredProducts = append(filteredProducts, product)
			}
		}
		return filteredProducts, nil
	}
	return products, nil
}