package abstract

import (
	"backend/src/internal/db/abstract"
	"backend/src/internal/domain"
)

// изм
type IBonusProgramInfoRepository interface {
	Create(conn abstract.IDBConnection, bonusProgramInfo *domain.BonusProgramInfo) error
	Delete(conn abstract.IDBConnection, programInfoID int) error
	GetByProgramID(conn abstract.IDBConnection, programID int) (*domain.BonusProgramInfo, error)
}
