package model

import (
	"backend/src/internal/domain"
	"backend/src/pkg/password"
	"time"
)

type Client struct {
	ID           int       `gorm:"column:client_id;primaryKey"`
	PasswordHash string    `gorm:"column:password_hash;not null"`
	FirstName    string    `gorm:"column:first_name;not null"`
	MiddleName   *string   `gorm:"column:middle_name"`
	LastName     string    `gorm:"column:last_name;not null"`
	PhoneNumber  string    `gorm:"column:phone_number;uniqueIndex;not null"`
	Email        string    `gorm:"column:email;uniqueIndex;not null"`
	Birthday     time.Time `gorm:"column:birthday;not null"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (modelObj *Client) ToDomain() (*domain.Client, error) {
	return ToDomain[Client, domain.Client](modelObj)
}

func (modelObj *Client) ToModel(domainObj *domain.Client) (*Client, error) {
	model, err := ToModel[Client, domain.Client](domainObj)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := password.Hash(domainObj.Password)
	if err != nil {
		return nil, err
	}

	model.PasswordHash = hashedPassword
	return model, nil
}
