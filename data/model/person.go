package model

import "time"

// PersonStatus type
type PersonStatus int

const (
	// StatusPending status
	StatusPending PersonStatus = iota
	// StatusRegistered status
	StatusRegistered
)

// Person DTO
type Person struct {
	ID             Key          `json:"id"`
	FirstName      string       `json:"first_name"`
	LastName       string       `json:"last_name"`
	Birthdate      time.Time    `json:"birthdate"`
	PassportNumber string       `json:"passport_number"`
	CountryOrigin  string       `json:"country_origin"`
	CountryBirth   string       `json:"country_birth"`
	Language       string       `json:"language"`
	Gender         string       `json:"gender"`
	Transafer      string       `json:"transafer"`
	MasterCouncil  string       `json:"mastercouncil"`
	Image          string       `json:"image"`
	Status         PersonStatus `json:"status"`
	Section        string       `json:"section"`
	Position       string       `json:"position"`
	Notes          string       `json:"notes"`
	UpdatedBy      Key          `json:"updated_by"`
}
