package impl

import (
	"backend/src/internal/db/abstract"
	"backend/src/internal/domain"
	"backend/src/internal/model"

	"gorm.io/gorm"
)

type ClientBonusProgramRepository struct{}

func NewClientBonusProgramRepository() *ClientBonusProgramRepository {
	return &ClientBonusProgramRepository{}
}

func (cbp *ClientBonusProgramRepository) Upsert(conn abstract.IDBConnection, bonus *domain.ClientBonusProgram) error {
	db := conn.Get().(*gorm.DB)

	bonusDAO := &model.ClientBonusProgram{}
	bonusDAO, err := bonusDAO.ToModel(bonus)
	if err != nil {
		return err
	}

	return db.Create(bonusDAO).Error
}

func (cbp *ClientBonusProgramRepository) Delete(conn abstract.IDBConnection, bonusID int) error {
	db := conn.Get().(*gorm.DB)
	return db.Where("client_bonus_program_id = ?", bonusID).Delete(&model.ClientBonusProgram{}).Error
}
