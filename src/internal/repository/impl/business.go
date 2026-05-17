package impl

import (
	"backend/src/internal/db/abstract"
	"backend/src/internal/domain"
	"backend/src/internal/model"

	"gorm.io/gorm"
)

type BusinessRepository struct{}

func NewBusinessRepository() *BusinessRepository {
	return &BusinessRepository{}
}

func (r *BusinessRepository) Create(conn abstract.IDBConnection, business *domain.Business) error {
	db := conn.Get().(*gorm.DB)

	businessDAO := &model.Business{}
	businessDAO, err := businessDAO.ToModel(business)
	if err != nil {
		return err
	}
	return db.Save(businessDAO).Error
}

func (r *BusinessRepository) Delete(conn abstract.IDBConnection, businessID int) error {
	db := conn.Get().(*gorm.DB)
	return db.Where("business_id = ?", businessID).Delete(&model.Business{}).Error
}

func (r *BusinessRepository) GetByOwnerID(conn abstract.IDBConnection, ownerID int) (*domain.Business, error) {
	db := conn.Get().(*gorm.DB)

	var ownerDAO model.Business
	err := db.Where("owner_id = ?", ownerID).First(&ownerDAO).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return ownerDAO.ToDomain()
}

// изм
