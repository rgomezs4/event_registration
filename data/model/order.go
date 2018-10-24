package model

// OrderHeader DTO
type OrderHeader struct {
	ID            Key           `json:"id"`
	PersonID      Key           `json:"person_id"`
	FirstName     string        `json:"first_name"`
	LastName      string        `json:"last_name"`
	PaymentMethod int           `json:"payment_method"`
	Total         float64       `json:"total"`
	Comment       string        `json:"comment"`
	BtcAddress    string        `json:"btc_address"`
	CreatedBy     Key           `json:"created_by"`
	Detail        []OrderDetail `json:"detail,omitempty"`
}

// OrderDetail DTO
type OrderDetail struct {
	ID            Key     `json:"id"`
	OrderHeaderID Key     `json:"order_enc_id"`
	ProductID     Key     `json:"product_id"`
	Name          string  `json:"name"`
	Size          string  `json:"size"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
	Amount        float64 `json:"amount"`
}
