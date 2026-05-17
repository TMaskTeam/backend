package abstract

import (
	"backend/src/internal/db/abstract"
	"backend/src/internal/domain"
)

// изм
type IBonusProgramRepository interface {
	Create(conn abstract.IDBConnection, bonusProgram *domain.BonusProgram) error
	Delete(conn abstract.IDBConnection, programID int) error
	GetByBusinessID(conn abstract.IDBConnection, businessID int) (*domain.BonusProgram, error)
}
