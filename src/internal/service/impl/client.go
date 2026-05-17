package impl

import (
	connection "backend/src/internal/db/abstract"
	"backend/src/internal/domain"
	repository "backend/src/internal/repository/abstract"
	"backend/src/pkg/jwt"
	"backend/src/pkg/password"
	"errors"
	"time"
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

func (s *ClientService) Login(login, pw string) (string, time.Time, *domain.Client, error) {
	client, err := s.clientRepo.GetByLogin(s.conn, login)
	if err != nil {
		return "", time.Time{}, nil, err
	}
	if client == nil {
		return "", time.Time{}, nil, errors.New("this login does not exists")
	}

	hash, err := s.clientRepo.GetPasswordHashById(s.conn, client.ID)
	if err != nil {
		return "", time.Time{}, nil, err
	}

	if err := password.CheckHash(hash, pw); err != nil {
		return "", time.Time{}, nil, errors.New("invalid credentials")
	}

	token, expiresAt, err := jwt.GenerateToken(client.ID, "business_owner")
	if err != nil {
		return "", time.Time{}, nil, err
	}

	return token, expiresAt, client, nil
}

func (cs *ClientService) Register(newClient *domain.Client) error {
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

func (cs *ClientService) GetByID(id int) (*domain.Client, error) {
	exists, err := cs.clientRepo.GetByID(cs.conn, id)
	if err != nil {
		return nil, err
	}
	if exists == nil {
		return nil, errors.New("user does not exist")
	}

	return exists, nil
}
