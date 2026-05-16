package dto

import "time"

type BusinessOwnerLoginRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type BusinessOwnerLoginResponse struct {
	Token     string                `json:"token"`
	ExpiresAt time.Time             `json:"expires_at"`
	Owner     BusinessOwnerResponse `json:"owner"`
}

type BusinessOwnerRegisterRequest struct {
	FirstName   string  `json:"first_name" validate:"required,min=2"`
	MiddleName  *string `json:"middle_name,omitempty"`
	LastName    string  `json:"last_name" validate:"required,min=2"`
	INN         string  `json:"inn" validate:"required,min=10,max=12"`
	PhoneNumber string  `json:"phone_number" validate:"required"`
	Email       string  `json:"email" validate:"required,email"`
	Birthday    string  `json:"birthday" validate:"required"`
	Password    string  `json:"password" validate:"required,min=8"`
}

type BusinessOwnerResponse struct {
	ID          int       `json:"owner_id"`
	FirstName   string    `json:"first_name"`
	MiddleName  *string   `json:"middle_name,omitempty"`
	LastName    string    `json:"last_name"`
	INN         string    `json:"inn"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Birthday    time.Time `json:"birthday"`
}
