package domain

import "time"

type BusinessOwner struct {
	ID           int
	PasswordHash string
	FirstName    string
	MiddleName   *string
	LastName     string
	INN          string
	Phone_number string
	Email        string
	Birthday     time.Time
}
