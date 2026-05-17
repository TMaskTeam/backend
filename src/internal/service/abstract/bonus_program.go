package abstract

import (
	"backend/src/internal/domain"
)

type IBonusProgramService interface {
	Create(businessID int, programName, tokenName string) (*domain.BonusProgram, error)
	GetByBusinessID(businessID int) ([]*domain.BonusProgram, error)
	GetAll() ([]*domain.BonusProgram, error)
	GetByProgramID(programID int) (*domain.BonusProgram, error)
	Update(programID int, programName, tokenName string) (*domain.BonusProgram, error)
	Delete(programID int) error
}
