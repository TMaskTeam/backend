package abstract

import (
	"backend/src/internal/db/abstract"
	"backend/src/internal/domain"
)

type IBusinessRepository interface {
	Create(conn abstract.IDBConnection, business *domain.Business) error
	Delete(conn abstract.IDBConnection, businessID int) error

	GetByOwnerID(conn abstract.IDBConnection, ownerID int) ([]domain.Business, error)
	GetByBusinessID(conn abstract.IDBConnection, businessID int) (*domain.Business, error)

	Update(conn abstract.IDBConnection, business *domain.Business) error
}
