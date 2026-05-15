package domain

import "time"

type Client struct {
	ID          int
	FirstName   string
	MiddleName  *string
	PhoneNumber string
	LastName    string
	Password    string
	Email       string
	Birthday    time.Time
}
