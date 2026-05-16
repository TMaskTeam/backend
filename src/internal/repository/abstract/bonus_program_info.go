package abstract

import (
	"backend/src/internal/db/abstract"
	"backend/src/internal/domain"
)

type IBonusProgramInfoRepository interface {
	Upsert(conn abstract.IDBConnection, bonus_program_info *domain.BonusProgramInfo) error

	Delete(conn abstract.IDBConnection, programInfoID int) error

	GetByID(conn abstract.IDBConnection, programInfoID int) (*domain.BonusProgramInfo, error)
}
