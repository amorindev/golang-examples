package domain

type Product struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Desc         string  `json:"desc"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	CategoryName string  `json:"category_name"`
}
