package impl

import (
	connection "backend/src/internal/db/abstract"
	"backend/src/internal/domain"
	repository "backend/src/internal/repository/abstract"
	"errors"
)

type ClientJoinService struct {
	conn                   connection.IDBConnection
	clientBonusProgramRepo repository.IClientBonusProgramRepository
}

func NewClientJoinService(
	conn connection.IDBConnection,
	clientBonusProgramRepo repository.IClientBonusProgramRepository,
) *ClientJoinService {
	return &ClientJoinService{
		conn:                   conn,
		clientBonusProgramRepo: clientBonusProgramRepo,
	}
}

func (cjs *ClientJoinService) JoinProgram(clientID int, programID int) (*int, error) {
	clientBonusProgram, err := cjs.clientBonusProgramRepo.GetByClientID(cjs.conn, clientID)
	if err != nil {
		if clientBonusProgram == nil {
			return nil, errors.New("client already in bonus program")
		}
		return nil, err
	}

	newClientBonusProgram := &domain.ClientBonusProgram{
		ProgramID:   programID,
		ClientID:    clientID,
		TokensCount: 0,
	}
	err = cjs.clientBonusProgramRepo.Upsert(cjs.conn, newClientBonusProgram)
	if err != nil {
		return nil, err
	}
	return &newClientBonusProgram.ID, nil
}
