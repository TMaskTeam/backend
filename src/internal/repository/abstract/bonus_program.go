package abstract

import (
	"backend/src/internal/db/abstract"
	"backend/src/internal/domain"
)

type IBonusProgramRepository interface {
	Create(conn abstract.IDBConnection, bonusProgram *domain.BonusProgram) error
	Delete(conn abstract.IDBConnection, programID int) error
	GetByBusinessID(conn abstract.IDBConnection, businessID int) ([]*domain.BonusProgram, error)
	GetAll(conn abstract.IDBConnection) ([]*domain.BonusProgram, error)
	GetByProgramID(conn abstract.IDBConnection, programID int) (*domain.BonusProgram, error)
	Update(conn abstract.IDBConnection, program *domain.BonusProgram) error
}
