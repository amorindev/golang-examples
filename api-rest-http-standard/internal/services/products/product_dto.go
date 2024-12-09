package products

import (
	"strings"
)

var products []Product

func getProductsDto(category string) ([]Product, error) {
	if strings.Compare(category,"Technology") == 0 {
		var filteredProducts []Product
		for _, product := range products {
			if strings.Compare(product.Category,category) == 0 {
				filteredProducts = append(filteredProducts, product)
			}
		}
		return filteredProducts, nil
	}
	return products, nil
}

func getProductDto(id int) (Product, error) {
	for _, product := range products {
		if(product.Id == id) {
			return product, nil
		}
	}
	return Product{}, productNotFound{}
}

func createProductDto(product Product) error {
	if len(products) == 0 {
		product.Id = 1
	} else {
		longitud := len(products)
		product.Id = products[longitud-1].Id + 1
	}

	products = append(products, product)
	return nil
}

func updateProductDto(id int, newProduct Product) error {
	for i, product := range products {
		if product.Id == id {
			newProduct.Id = id
			products[i] = newProduct
			return nil
		}
	}
	return productNotFound{}
}

func deleteProductDto(id int) error {
	for i, product := range products {
		if product.Id == id {
			products = append(products[:i], products[i+1:]...)
			return nil
		}
	}
	return productNotFound{}
}
