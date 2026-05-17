package dto

import "time"

type ClientJoinRequest struct {
	Client    ClientResponse `json:"client"`
	ProgramID int            `json:"program_id"`
}

type ClientJoinResponse struct {
	Client               ClientResponse `json:"client"`
	ClientBonusProgramID int            `json:"client_bonus_program_id"`
	ProgramID            int            `json:"program_id"`
	TokensCount          int            `json:"tokens_count"`
}

type ClientLoginRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ClientLoginResponse struct {
	Token     string         `json:"token"`
	ExpiresAt time.Time      `json:"expires_at"`
	Owner     ClientResponse `json:"client"`
}

type ClientRegisterRequest struct {
	FirstName   string  `json:"first_name" validate:"required,min=2"`
	MiddleName  *string `json:"middle_name,omitempty"`
	LastName    string  `json:"last_name" validate:"required,min=2"`
	PhoneNumber string  `json:"phone_number" validate:"required"`
	Email       string  `json:"email" validate:"required,email"`
	Birthday    string  `json:"birthday" validate:"required"`
	Password    string  `json:"password" validate:"required,min=8"`
}

type ClientResponse struct {
	ID          int       `json:"client_id"`
	FirstName   string    `json:"first_name"`
	MiddleName  *string   `json:"middle_name,omitempty"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Birthday    time.Time `json:"birthday"`
}
