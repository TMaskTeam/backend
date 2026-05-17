package abstract

import (
	"backend/src/internal/domain"
)

type IClientBonusProgramService interface {
	GetAllByClientID(clientID int) ([]*domain.ClientBonusProgram, error)
}
