package abstract

import (
	"backend/src/internal/db/abstract"
	"backend/src/internal/domain"
)

type IBusinessOwnerRepository interface {
	Upsert(conn abstract.IDBConnection, owner *domain.BusinessOwner) error
	Delete(conn abstract.IDBConnection, ownerID int) error

	GetByINN(conn abstract.IDBConnection, inn string) (*domain.BusinessOwner, error)
	GetByPhoneNumber(conn abstract.IDBConnection, phoneNumber string) (*domain.BusinessOwner, error)
	GetByEmail(conn abstract.IDBConnection, email string) (*domain.BusinessOwner, error)

	ExistsByINN(conn abstract.IDBConnection, inn string) (bool, error)
	ExistsByPhoneNumber(conn abstract.IDBConnection, phoneNumber string) (bool, error)
	ExistsByEmail(conn abstract.IDBConnection, email string) (bool, error)
}
