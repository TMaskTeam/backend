package abstract

import "backend/src/internal/domain"

type IClientService interface {
	Register(client *domain.Client) error
}
