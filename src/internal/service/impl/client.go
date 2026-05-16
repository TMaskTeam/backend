package impl

import (
	connection "backend/src/internal/db/abstract"
	"backend/src/internal/domain"
	repository "backend/src/internal/repository/abstract"
	"errors"
)

type ClientService struct {
	conn       connection.IDBConnection
	clientRepo repository.IClientRepository
}

func NewClientService(
	conn connection.IDBConnection,
	clientRepo repository.IClientRepository,
) *ClientService {
	return &ClientService{
		conn:       conn,
		clientRepo: clientRepo,
	}
}

func (cs *ClientService) RegisterClient(newClient *domain.Client) error {
	exists, err := cs.clientRepo.GetByEmail(cs.conn, newClient.Email)
	if err != nil {
		return err
	}
	if exists != nil {
		return errors.New("email is already used")
	}

	exists, err = cs.clientRepo.GetByPhoneNumber(cs.conn, newClient.PhoneNumber)
	if err != nil {
		return err
	}
	if exists != nil {
		return errors.New("phone number is already used")
	}

	err = cs.clientRepo.Upsert(cs.conn, newClient)
	if err != nil {
		return err
	}

	return nil
}
