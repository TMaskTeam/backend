package abstract

import "backend/src/internal/domain"

type IBusinessOwnerService interface {
	Register(owner *domain.BusinessOwner) (*domain.BusinessOwner, error)
}
