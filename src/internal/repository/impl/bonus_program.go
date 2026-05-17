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

func (r *BonusProgramRepository) Create(conn abstract.IDBConnection, bonusProgram *domain.BonusProgram) error {
	db := conn.Get().(*gorm.DB)

	bonusProgramDAO := &model.BonusProgram{}
	bonusProgramDAO, err := bonusProgramDAO.ToModel(bonusProgram)
	if err != nil {
		return err
	}

	if err := db.Create(bonusProgramDAO).Error; err != nil {
		return err
	}

	bonusProgram.ProgramID = bonusProgramDAO.ProgramID

	return nil
}

func (r *BonusProgramRepository) Delete(conn abstract.IDBConnection, programID int) error {
	db := conn.Get().(*gorm.DB)
	return db.Where("program_id = ?", programID).Delete(&model.BonusProgram{}).Error
}

func (r *BonusProgramRepository) GetByBusinessID(conn abstract.IDBConnection, businessID int) ([]*domain.BonusProgram, error) {
	db := conn.Get().(*gorm.DB)

	var bonusProgramsDAO []model.BonusProgram
	err := db.Where("business_id = ?", businessID).Find(&bonusProgramsDAO).Error
	if err != nil {
		return nil, err
	}

	result := make([]*domain.BonusProgram, 0, len(bonusProgramsDAO))
	for _, dao := range bonusProgramsDAO {
		domainObj, err := dao.ToDomain()
		if err != nil {
			return nil, err
		}
		result = append(result, domainObj)
	}

	return result, nil
}

func (r *BonusProgramRepository) GetAll(conn abstract.IDBConnection) ([]*domain.BonusProgram, error) {
	db := conn.Get().(*gorm.DB)

	var bonusProgramsDAO []model.BonusProgram
	err := db.Find(&bonusProgramsDAO).Error
	if err != nil {
		return nil, err
	}

	result := make([]*domain.BonusProgram, 0, len(bonusProgramsDAO))
	for _, dao := range bonusProgramsDAO {
		domainObj, err := dao.ToDomain()
		if err != nil {
			return nil, err
		}
		result = append(result, domainObj)
	}

	return result, nil
}
