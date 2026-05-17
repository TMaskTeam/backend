package abstract

import (
	"backend/src/internal/db/abstract"
	"backend/src/internal/domain"
)

type IClientRepository interface {
	Upsert(conn abstract.IDBConnection, client *domain.Client) error
	Delete(conn abstract.IDBConnection, clientID int) error

	GetPasswordHashById(conn abstract.IDBConnection, clientID int) (string, error)
	GetByPhoneNumber(conn abstract.IDBConnection, phoneNumber string) (*domain.Client, error)
	GetByEmail(conn abstract.IDBConnection, email string) (*domain.Client, error)

	GetByLogin(conn abstract.IDBConnection, login string) (*domain.Client, error)
}
