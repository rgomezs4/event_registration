package model

// JSONApiRequest - Basic structure of json api request
type JSONApiRequest struct {
	Data JSONApiRequestData `json:"data"`
}

// JSONApiRequestData - Basic structure of json api request data
type JSONApiRequestData struct {
	Type       string      `json:"type"`
	UserID     Key         `json:"user_id"`
	Attributes interface{} `json:"attributes"`
}

// Login structure for the request
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UpdateAdmin structure with only the fields needed to update
type UpdateAdmin struct {
	Name string `json:"name"`
}

// RegisterPerson structure for the request
type RegisterPerson struct {
	PersonID Key    `json:"id"`
	Image    string `json:"image"`
	Items    []Item `json:"items"`
}
