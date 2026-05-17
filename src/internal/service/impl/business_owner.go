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

func (s *BusinessOwnerService) Update(owner *domain.BusinessOwner) (*domain.BusinessOwner, error) {
	existing, err := s.ownerRepo.GetByID(s.conn, owner.ID)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, errors.New("user not found")
	}

	if owner.FirstName != "" {
		existing.FirstName = owner.FirstName
	}
	if owner.LastName != "" {
		existing.LastName = owner.LastName
	}
	if owner.MiddleName != nil {
		existing.MiddleName = owner.MiddleName
	}
	if owner.PhoneNumber != "" {
		existing.PhoneNumber = owner.PhoneNumber
	}
	if owner.Email != "" {
		existing.Email = owner.Email
	}
	if owner.Password != "" {
		hash, err := password.Hash(owner.Password)
		if err != nil {
			return nil, err
		}
		existing.Password = hash
	}

	return existing, s.ownerRepo.UpdateByID(s.conn, owner)
}

func (s *BusinessOwnerService) Login(login, pw string) (string, time.Time, *domain.BusinessOwner, error) {
	owner, err := s.ownerRepo.GetByLogin(s.conn, login)
	if err != nil {
		return "", time.Time{}, nil, err
	}
	if owner == nil {
		return "", time.Time{}, nil, errors.New("this login does not exists")
	}

	hash, err := s.ownerRepo.GetPasswordHashById(s.conn, owner.ID)
	if err != nil {
		return "", time.Time{}, nil, err
	}

	if err := password.CheckHash(hash, pw); err != nil {
		return "", time.Time{}, nil, errors.New("invalid credentials")
	}

	token, expiresAt, err := jwt.GenerateToken(owner.ID, "business_owner")
	if err != nil {
		return "", time.Time{}, nil, err
	}

	return token, expiresAt, owner, nil
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

func (s *BusinessOwnerService) GetByID(id int) (*domain.BusinessOwner, error) {
	exists, err := s.ownerRepo.GetByID(s.conn, id)
	if err != nil {
		return nil, err
	}
	if exists == nil {
		return nil, errors.New("user does not exist")
	}

	return exists, nil
}
