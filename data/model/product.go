package model

// Product DTO
type Product struct {
	ID      Key     `json:"id"`
	Name    string  `json:"name"`
	Barcode string  `json:"barcode"`
	Size    string  `json:"size"`
	Stock   int     `json:"stock"`
	Price   float64 `json:"price"`
}
