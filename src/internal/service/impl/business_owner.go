package impl

import (
	connection "backend/src/internal/db/abstract"
	"backend/src/internal/domain"
	repository "backend/src/internal/repository/abstract"
	"errors"
	"time"
)

type BusinessOwnerService struct {
	conn      connection.IDBConnection
	ownerRepo repository.IBusinessOwnerRepository
}

func NewBusinessOwnerService(
	conn connection.IDBConnection,
	ownerRepo repository.IBusinessOwnerRepository,
) *BusinessOwnerService {
	return &BusinessOwnerService{
		conn:      conn,
		ownerRepo: ownerRepo,
	}
}

func (s *BusinessOwnerService) Login(login, password string) (string, time.Time, *domain.BusinessOwner, error) {
	exists, err := s.ownerRepo.GetByLogin(s.conn, login)
	if err != nil {
		return "", time.Time{}, nil, err
	}
	if exists == nil {
		return "", time.Time{}, nil, errors.New("this login does not exists")
	}

}

func (s *BusinessOwnerService) Register(owner *domain.BusinessOwner) error {
	exists, err := s.ownerRepo.GetByINN(s.conn, owner.INN)
	if err != nil {
		return err
	}
	if exists != nil {
		return errors.New("inn is already used")
	}

	exists, err = s.ownerRepo.GetByEmail(s.conn, owner.Email)
	if err != nil {
		return err
	}
	if exists != nil {
		return errors.New("email is already used")
	}

	exists, err = s.ownerRepo.GetByPhoneNumber(s.conn, owner.PhoneNumber)
	if err != nil {
		return err
	}
	if exists != nil {
		return errors.New("phone number is already used")
	}

	err = s.ownerRepo.Upsert(s.conn, owner)
	if err != nil {
		return err
	}

	return nil
}
