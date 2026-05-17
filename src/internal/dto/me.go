package dto

import "time"

type BusinessOwnerMeResponse struct {
	Role        string    `json:"role"`
	ID          int       `json:"owner_id"`
	FirstName   string    `json:"first_name"`
	MiddleName  *string   `json:"middle_name,omitempty"`
	LastName    string    `json:"last_name"`
	INN         string    `json:"inn"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Birthday    time.Time `json:"birthday"`
}

type ClientMeResponse struct {
	Role        string    `json:"role"`
	ID          int       `json:"client_id"`
	FirstName   string    `json:"first_name"`
	MiddleName  *string   `json:"middle_name,omitempty"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Birthday    time.Time `json:"birthday"`
}

type UpdateProfileRequest struct {
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	MiddleName  string `json:"middle_name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
}
