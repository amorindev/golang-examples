package products

type productNotFound struct {}

func (p productNotFound) Error() string {
	return "Product not found"
}





