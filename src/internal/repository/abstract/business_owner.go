package abstract

import (
	"backend/src/internal/db/abstract"
	"backend/src/internal/domain"
)

type IBusinessOwnerRepository interface {
	UpdateByID(conn abstract.IDBConnection, owner *domain.BusinessOwner) error
	Upsert(conn abstract.IDBConnection, owner *domain.BusinessOwner) error
	Delete(conn abstract.IDBConnection, ownerID int) error

	GetPasswordHashById(conn abstract.IDBConnection, ownerID int) (string, error)
	GetByINN(conn abstract.IDBConnection, inn string) (*domain.BusinessOwner, error)
	GetByPhoneNumber(conn abstract.IDBConnection, phoneNumber string) (*domain.BusinessOwner, error)
	GetByEmail(conn abstract.IDBConnection, email string) (*domain.BusinessOwner, error)

	GetByLogin(conn abstract.IDBConnection, login string) (*domain.BusinessOwner, error)
	GetByID(conn abstract.IDBConnection, id int) (*domain.BusinessOwner, error)
}
