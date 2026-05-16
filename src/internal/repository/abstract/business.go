package abstract

import (
	"backend/src/internal/db/abstract"
	"backend/src/internal/domain"
)

type IBusinessRepository interface {
	Upsert(conn abstract.IDBConnection, business *domain.Business) error

	Delete(conn abstract.IDBConnection, business_id int) error

	GetByID(conn abstract.IDBConnection, business_id int) error
}
