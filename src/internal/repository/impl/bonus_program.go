package impl

import (
	"backend/src/internal/db/abstract"
	"backend/src/internal/domain"
	"backend/src/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BonusProgramRepository struct{}

func NewBonusProgramRepository() *BonusProgramRepository {
	return &BonusProgramRepository{}
}

func (r *BonusProgramRepository) Upsert(conn abstract.IDBConnection, bonus_program *domain.BonusProgram) error {
	db := conn.Get().(*gorm.DB)

	bonusProgramDAO := &model.BonusProgram{}
	bonusProgramDAO, err := bonusProgramDAO.ToModel(bonus_program)
	if err != nil {
		return err
	}
	return db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "program_id"},
		},
		UpdateAll: true,
	}).Create(bonusProgramDAO).Error
}

func (r *BonusProgramRepository) Delete(conn abstract.IDBConnection, programID int) error {
	db := conn.Get().(*gorm.DB)
	return db.Where("program_id = ?", programID).Delete(&model.BonusProgram{}).Error
}

func (r *BonusProgramRepository) GetByID(conn abstract.IDBConnection, programID int) (*domain.BonusProgram, error) {
	db := conn.Get().(*gorm.DB)

	var bonusProgramDAO model.BonusProgram
	err := db.Where("program_id = ?", programID).First(&bonusProgramDAO).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return bonusProgramDAO.ToDomain()
}
