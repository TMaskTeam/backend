package dto

import "time"

type ClientRegisterRequest struct {
	FirstName   string  `json:"first_name" validate:"required,min=2"`
	MiddleName  *string `json:"middle_name,omitempty"`
	LastName    string  `json:"last_name" validate:"required,min=2"`
	PhoneNumber string  `json:"phone_number" validate:"required"`
	Email       string  `json:"email" validate:"required,email"`
	Birthday    string  `json:"birthday" validate:"required"`
	Password    string  `json:"password" validate:"required,min=8"`
}

type ClientRegisterResponse struct {
	ID          int       `json:"owner_id"`
	FirstName   string    `json:"first_name"`
	MiddleName  *string   `json:"middle_name,omitempty"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Birthday    time.Time `json:"birthday"`
}
