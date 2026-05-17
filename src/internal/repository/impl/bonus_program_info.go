package impl

import (
	"backend/src/internal/db/abstract"
	"backend/src/internal/domain"
	"backend/src/internal/model"

	"gorm.io/gorm"
)

type BonusProgramInfoRepository struct{}

func NewBonusProgramInfoRepository() *BonusProgramInfoRepository {
	return &BonusProgramInfoRepository{}
}

func (r *BonusProgramInfoRepository) Create(conn abstract.IDBConnection, bonusProgramInfo *domain.BonusProgramInfo) error {
	db := conn.Get().(*gorm.DB)

	bonusProgramInfoDAO := &model.BonusProgramInfo{}
	bonusProgramInfoDAO, err := bonusProgramInfoDAO.ToModel(bonusProgramInfo)
	if err != nil {
		return err
	}
	return db.Save(bonusProgramInfoDAO).Error
}

func (r *BonusProgramInfoRepository) Delete(conn abstract.IDBConnection, programInfoID int) error {
	db := conn.Get().(*gorm.DB)
	return db.Where("program_info_id = ?", programInfoID).Delete(&model.BonusProgramInfo{}).Error
}

func (r *BonusProgramInfoRepository) GetByProgramID(conn abstract.IDBConnection, programID int) (*domain.BonusProgramInfo, error) {
	db := conn.Get().(*gorm.DB)

	var bonusProgramDAO model.BonusProgramInfo
	err := db.Where("program_id = ?", programID).First(&bonusProgramDAO).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return bonusProgramDAO.ToDomain()
}

// изм
