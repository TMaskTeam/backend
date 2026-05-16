package abstract

import (
	"backend/src/internal/db/abstract"
	"backend/src/internal/domain"
)

type IClientBonusProgramRepository interface {
	Upsert(conn abstract.IDBConnection, bonus *domain.ClientBonusProgram) error
	Delete(conn abstract.IDBConnection, bonusID int) error
	GetByClientID(conn abstract.IDBConnection, clientID int) error
	GetByProgramID(conn abstract.IDBConnection, programID int) error
}
