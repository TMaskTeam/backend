package domain

import "time"

type BusinessOwner struct {
	ID           int
	FirstName    string
	MiddleName   *string
	LastName     string
	INN          string
	phone_number string
	email        string
	birthday     time.Time
}
