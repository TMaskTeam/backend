package impl

import (
	connection "backend/src/internal/db/abstract"
	"backend/src/internal/domain"
	repository "backend/src/internal/repository/abstract"
)

type BonusProgramService struct {
	conn             connection.IDBConnection
	bonusProgramRepo repository.IBonusProgramRepository
}

func NewBonusProgramService(
	conn connection.IDBConnection,
	bonusProgramRepo repository.IBonusProgramRepository,
) *BonusProgramService {
	return &BonusProgramService{
		conn:             conn,
		bonusProgramRepo: bonusProgramRepo,
	}
}

func (s *BonusProgramService) Create(businessID int, programName, tokenName string) (*domain.BonusProgram, error) {
	program := &domain.BonusProgram{
		BusinessID:  businessID,
		ProgramName: programName,
		TokenName:   tokenName,
	}

	err := s.bonusProgramRepo.Create(s.conn, program)
	if err != nil {
		return nil, err
	}

	return program, nil
}

func (s *BonusProgramService) GetByBusinessID(businessID int) ([]*domain.BonusProgram, error) {
	return s.bonusProgramRepo.GetByBusinessID(s.conn, businessID)
}

func (s *BonusProgramService) GetAll() ([]*domain.BonusProgram, error) {
	return s.bonusProgramRepo.GetAll(s.conn)
}
