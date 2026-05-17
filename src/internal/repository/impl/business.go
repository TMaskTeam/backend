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

func (r *BusinessRepository) GetByOwnerID(conn abstract.IDBConnection, ownerID int) ([]domain.Business, error) {
	db := conn.Get().(*gorm.DB)

	var ownerDAOs []model.Business
	err := db.Where("owner_id = ?", ownerID).Find(&ownerDAOs).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return (&model.Business{}).ToDomainSlice(ownerDAOs)
}

func (r *BusinessRepository) GetByBusinessID(conn abstract.IDBConnection, businessID int) (*domain.Business, error) {
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
