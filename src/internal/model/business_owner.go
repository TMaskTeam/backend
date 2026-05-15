package model

import "time"

type BusinessOwner struct {
	ID           int       `gorm:"column:owner_id;primaryKey"`
	PasswordHash string    `gorm:"column:password_hash;not null"`
	FirstName    string    `gorm:"column:first_name;not null"`
	MiddleName   *string   `gorm:"column:middle_name"`
	LastName     string    `gorm:"column:last_name;not null"`
	INN          string    `gorm:"column:inn;uniqueIndex;not null"`
	PhoneNumber  string    `gorm:"column:phone_number;uniqueIndex;not null"`
	Email        string    `gorm:"column:email;uniqueIndex;not null"`
	Birthday     time.Time `gorm:"column:birthday;not null"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
