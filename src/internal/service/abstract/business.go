package abstract

import "backend/src/internal/domain"

type IBusinessService interface {
	GetByOwnerID(ownerID int) ([]domain.Business, error)
	Delete(businessID, ownerID int) error
	Create(ownerID int, name string, address string) (*domain.Business, error)
}
