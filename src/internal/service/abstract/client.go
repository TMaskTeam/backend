package abstract

import "backend/src/internal/domain"

type IClientService interface {
	RegisterClient(client *domain.Client) error
}
