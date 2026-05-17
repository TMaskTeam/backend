package abstract

import (
	"backend/src/internal/domain"
	"time"
)

type IBusinessOwnerService interface {
	Login(login, password string) (string, time.Time, *domain.BusinessOwner, error)
	Register(owner *domain.BusinessOwner) error

	GetByID(id int) (*domain.BusinessOwner, error)
}
