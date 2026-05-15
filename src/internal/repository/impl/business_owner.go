package impl

import (
	"backend/src/internal/db/abstract"
	"backend/src/internal/domain"
)

type BusinessOwnerRepository struct{}

func NewBusinessOwnerRepository() BusinessOwnerRepository {
	return BusinessOwnerRepository{}
}

func (bo *BusinessOwnerRepository) Upsert(conn abstract.IDBConnection, owner *domain.BusinessOwner) error {

}
