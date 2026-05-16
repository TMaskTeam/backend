package impl

import (
	"backend/src/internal/db/abstract"
	"backend/src/internal/domain"
	"backend/src/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BonusProgramInfoRepository struct{}

func NewBonusProgramInfoRepository() *BonusProgramInfoRepository {
	return &BonusProgramInfoRepository{}
}

func (r *BonusProgramInfoRepository) Upsert(conn abstract.IDBConnection, bonus_program_info *domain.BonusProgramInfo) error {
	db := conn.Get().(*gorm.DB)

	bonusProgramInfoDAO := &model.BonusProgramInfo{}
	bonusProgramInfoDAO, err := bonusProgramInfoDAO.ToModel(bonus_program_info)
	if err != nil {
		return err
	}
	return db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "program_info_id"},
		},
		UpdateAll: true,
	}).Create(bonusProgramInfoDAO).Error
}

func (r *BonusProgramInfoRepository) Delete(conn abstract.IDBConnection, programInfoID int) error {
	db := conn.Get().(*gorm.DB)
	return db.Where("program_info_id = ?", programInfoID).Delete(&model.BonusProgramInfo{}).Error
}

func (r *BonusProgramInfoRepository) GetByID(conn abstract.IDBConnection, programInfoID int) (*domain.BonusProgramInfo, error) {
	db := conn.Get().(*gorm.DB)

	var bonusProgramInfoDAO model.BonusProgramInfo
	err := db.Where("program_info_id = ?", programInfoID).First(&bonusProgramInfoDAO).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return bonusProgramInfoDAO.ToDomain()
}
