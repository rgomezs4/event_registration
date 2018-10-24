package model

// Admin DTO
type Admin struct {
	ID       Key    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password,omitempty"`
}
