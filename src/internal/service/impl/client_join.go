package impl

import (
	connection "backend/src/internal/db/abstract"
	"backend/src/internal/domain"
	repository "backend/src/internal/repository/abstract"
	"backend/src/pkg/password"
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

func (cs *ClientJoinService) Join(client *domain.Client, programID int) (int, int, error) {
	hash, err := cs.clientBonusProgramRepo.GetByProgramID(cs.conn, client.ID)
	if err != nil {
		return -1, -1, errors.New("this client does not exists")
	}

	if err := password.CheckHash(hash, client.Password); err != nil {
		return -1, -1, errors.New("invalid credentials")
	}

	existClientBonusProgram, err := cs.clientBonusProgramRepo.GetByClientID(cs.conn, client.ID)
	if existClientBonusProgram != nil {
		return existClientBonusProgram.ID, existClientBonusProgram.TokensCount, nil
	}

	newClientBonusProgram := &domain.ClientBonusProgram{}
	err = cs.clientBonusProgramRepo.Upsert(cs.conn, newClientBonusProgram)
	if err != nil {
		return -1, -1, err
	}

	return newClientBonusProgram.ID, 0, nil
}
