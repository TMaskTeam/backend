package dto

import "time"

type BusinessOwnerRegisterRequest struct {
	FirstName   string    `json:"first_name"`
	MiddleName  *string   `json:"middle_name,omitempty"`
	LastName    string    `json:"last_name"`
	INN         string    `json:"inn"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Birthday    time.Time `json:"birthday"`
	Password    string    `json:"password"`
}

type BusinessOwnerRegisterResponse struct {
	ID          int       `json:"owner_id"`
	FirstName   string    `json:"first_name"`
	MiddleName  *string   `json:"middle_name,omitempty"`
	LastName    string    `json:"last_name"`
	INN         string    `json:"inn"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Birthday    time.Time `json:"birthday"`
}
