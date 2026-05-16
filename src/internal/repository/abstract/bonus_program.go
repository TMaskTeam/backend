package abstract

import (
	"backend/src/internal/db/abstract"
	"backend/src/internal/domain"
)

type IBonusProgramRepository interface {
	Upsert(conn abstract.IDBConnection, bonus_program *domain.BonusProgram) error

	Delete(conn abstract.IDBConnection, programID int) error

	GetByID(conn abstract.IDBConnection, programID int) (*domain.BonusProgram, error)
}
