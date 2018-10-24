package model

// Item DTO
type Item struct {
	ID   Key    `json:"id"`
	Name string `json:"name"`
}

// PersonItem DTO
type PersonItem struct {
	ID        Key    `json:"id"`
	PersonID  Key    `json:"person_id"`
	ItemID    Key    `json:"item_id"`
	ItemName  string `json:"item_name"`
	CreatedBy Key    `json:"created_by"`
}
