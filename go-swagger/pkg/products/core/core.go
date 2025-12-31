package core

// ProductCreateReq is a struct that represents the parameters needed to create a product.
type ProductCreateReq struct {
    Name         string  `json:"name"`
	Desc         string  `json:"desc"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	CategoryName string  `json:"category_name"`
}

