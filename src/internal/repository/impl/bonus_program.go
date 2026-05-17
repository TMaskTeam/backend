package impl

import (
	"backend/src/internal/db/abstract"
	"backend/src/internal/domain"
	"backend/src/internal/model"

	"gorm.io/gorm"
)

type BonusProgramRepository struct{}

func NewBonusProgramRepository() *BonusProgramRepository {
	return &BonusProgramRepository{}
}

func (r *BonusProgramRepository) Upsert(conn abstract.IDBConnection, bonusProgram *domain.BonusProgram) error {
	db := conn.Get().(*gorm.DB)

	bonusProgramDAO := &model.BonusProgram{}
	bonusProgramDAO, err := bonusProgramDAO.ToModel(bonusProgram)
	if err != nil {
		return err
	}
	return db.Save(bonusProgramDAO).Error
}

func (r *BonusProgramRepository) Delete(conn abstract.IDBConnection, programID int) error {
	db := conn.Get().(*gorm.DB)
	return db.Where("program_id = ?", programID).Delete(&model.BonusProgram{}).Error
}

func (r *BonusProgramRepository) GetByBusinessID(conn abstract.IDBConnection, businessID int) (*domain.BonusProgram, error) {
	db := conn.Get().(*gorm.DB)

	var bonusProgramDAO model.BonusProgram
	err := db.Where("business_id = ?", businessID).First(&bonusProgramDAO).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return bonusProgramDAO.ToDomain()
}
