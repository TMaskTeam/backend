package impl

import (
	"backend/src/internal/db/abstract"
	"backend/src/internal/domain"
	"backend/src/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BusinessOwnerRepository struct{}

func NewBusinessOwnerRepository() BusinessOwnerRepository {
	return BusinessOwnerRepository{}
}

func (bo *BusinessOwnerRepository) Upsert(conn abstract.IDBConnection, owner *domain.BusinessOwner) error {
	db := conn.Get().(*gorm.DB)

	ownerDAO := &model.BusinessOwner{}
	ownerDAO, err := ownerDAO.ToModel(owner)
	if err != nil {
		return err
	}

	return db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "code"},
			{Name: "owner_id"},
			{Name: "phone_number"},
			{Name: "inn"},
			{Name: "email"},
		},
		UpdateAll: true,
	}).
		Create(ownerDAO).Error
}

func (bo *BusinessOwnerRepository) Delete(conn abstract.IDBConnection, ownerID int) error {
	db := conn.Get().(*gorm.DB)
	return db.Where("owner_id = ?", ownerID).Delete(&model.BusinessOwner{}).Error
}

func (bo *BusinessOwnerRepository) GetPasswordHashById(conn abstract.IDBConnection, ownerID int) (string, error) {
	db := conn.Get().(*gorm.DB)

	var hashedPassword string
	err := db.Model(&model.BusinessOwner{}).
		Where("owner_id = ?", ownerID).
		Select("password_hash").Error
	if err != nil {
		return "", err
	}

	return hashedPassword, err
}

func (bo *BusinessOwnerRepository) GetByINN(conn abstract.IDBConnection, inn string) (*domain.BusinessOwner, error) {
	db := conn.Get().(*gorm.DB)

	var ownerDAO model.BusinessOwner
	err := db.Where("inn = ?", inn).First(&ownerDAO).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return ownerDAO.ToDomain()
}

func (bo *BusinessOwnerRepository) GetByPhoneNumber(conn abstract.IDBConnection, phoneNumber string) (*domain.BusinessOwner, error) {
	db := conn.Get().(*gorm.DB)

	var ownerDAO model.BusinessOwner
	err := db.Where("phone_number = ?", phoneNumber).First(&ownerDAO).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return ownerDAO.ToDomain()
}

func (bo *BusinessOwnerRepository) GetByEmail(conn abstract.IDBConnection, email string) (*domain.BusinessOwner, error) {
	db := conn.Get().(*gorm.DB)

	var ownerDAO model.BusinessOwner
	err := db.Where("email = ?", email).First(&ownerDAO).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return ownerDAO.ToDomain()
}
