package models

type Person struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Address   *Address
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}
