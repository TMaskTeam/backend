package domain

import "time"

type BusinessOwner struct {
	ID          int
	Password    string
	FirstName   string
	MiddleName  *string
	LastName    string
	INN         string
	PhoneNumber string
	Email       string
	Birthday    time.Time
}
