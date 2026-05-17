package abstract

import (
	"backend/src/internal/domain"
	"time"
)

type IClientService interface {
	Login(login, password string) (string, time.Time, *domain.Client, error)
	Register(client *domain.Client) error

	GetByID(id int) (*domain.Client, error)
	Update(client *domain.Client) error
}
