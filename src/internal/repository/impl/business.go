package impl

import (
	"backend/src/internal/db/abstract"
	"backend/src/internal/domain"
	"backend/src/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BusinessRepository struct{}

func NewBusinessRepository() *BusinessRepository {
	return &BusinessRepository{}
}

func (r *BusinessRepository) Upsert(conn abstract.IDBConnection, business *domain.Business) error {
	db := conn.Get().(*gorm.DB)

	businessDAO := &model.Business{}
	businessDAO, err := businessDAO.ToModel(business)
	if err != nil {
		return err
	}
	return db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "business_id"},
		},
		UpdateAll: true,
	}).Create(businessDAO).Error
}

func (r *BusinessRepository) Delete(conn abstract.IDBConnection, businessID int) error {
	db := conn.Get().(*gorm.DB)
	return db.Where("business_id = ?", businessID).Delete(&model.Business{}).Error
}

func (r *BusinessRepository) GetByID(conn abstract.IDBConnection, businessID int) (*domain.Business, error) {
	db := conn.Get().(*gorm.DB)

	var businessDAO model.Business
	err := db.Where("business_id = ?", businessID).First(&businessDAO).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return businessDAO.ToDomain()
}
