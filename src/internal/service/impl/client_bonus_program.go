package impl

import (
	connection "backend/src/internal/db/abstract"
	"backend/src/internal/domain"
	repository "backend/src/internal/repository/abstract"
)

type ClientBonusProgramService struct {
	conn                   connection.IDBConnection
	clientbonusProgramRepo repository.IClientBonusProgramRepository
}

func (cbps *ClientBonusProgramService) GetAllByClientID(clientID int) ([]*domain.ClientBonusProgram, error) {
	return cbps.clientbonusProgramRepo.GetAllWithClientID(cbps.conn, clientID)
}
