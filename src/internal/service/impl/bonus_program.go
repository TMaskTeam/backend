package impl

import (
	connection "backend/src/internal/db/abstract"
	"backend/src/internal/domain"
	repository "backend/src/internal/repository/abstract"
	"errors"
)

type BonusProgramService struct {
	conn             connection.IDBConnection
	businessRepo     repository.IBusinessRepository
	bonusProgramRepo repository.IBonusProgramRepository
}

func NewBonusProgramService(
	conn connection.IDBConnection,
	businessRepo repository.IBusinessRepository,
	bonusProgramRepo repository.IBonusProgramRepository,
) *BonusProgramService {
	return &BonusProgramService{
		conn:             conn,
		bonusProgramRepo: bonusProgramRepo,
		businessRepo:     businessRepo,
	}
}

func (s *BonusProgramService) Create(businessID, ownerID int, programName, tokenName string) (*domain.BonusProgram, error) {
	business, err := s.businessRepo.GetByID(s.conn, businessID)
	if err != nil {
		return nil, err
	}
	if business == nil {
		return nil, errors.New("business not found")
	}
	if business.OwnerID != ownerID {
		return nil, errors.New("you don't have permission to create program for this business")
	}

	program := &domain.BonusProgram{
		BusinessID:  businessID,
		ProgramName: programName,
		TokenName:   tokenName,
	}

	err = s.bonusProgramRepo.Create(s.conn, program)
	return program, err
}

func (s *BonusProgramService) GetByBusinessID(businessID int) ([]*domain.BonusProgram, error) {
	return s.bonusProgramRepo.GetByBusinessID(s.conn, businessID)
}

func (s *BonusProgramService) GetAll() ([]*domain.BonusProgram, error) {
	return s.bonusProgramRepo.GetAll(s.conn)
}
