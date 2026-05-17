package impl

import (
	"backend/src/internal/db/abstract"
	"backend/src/internal/domain"
	"backend/src/internal/model"
	"time"

	"gorm.io/gorm"
)

type BusinessOwnerRepository struct{}

func NewBusinessOwnerRepository() *BusinessOwnerRepository {
	return &BusinessOwnerRepository{}
}

func (bo *BusinessOwnerRepository) UpdateByID(conn abstract.IDBConnection, owner *domain.BusinessOwner) error {
	db := conn.Get().(*gorm.DB)

	ownerDAO := &model.BusinessOwner{}
	ownerDAO, err := ownerDAO.ToModel(owner)
	if err != nil {
		return err
	}

	return db.Model(&model.BusinessOwner{}).
		Where("owner_id = ?", owner.ID).
		Updates(map[string]interface{}{
			"first_name":    ownerDAO.FirstName,
			"last_name":     ownerDAO.LastName,
			"middle_name":   ownerDAO.MiddleName,
			"phone_number":  ownerDAO.PhoneNumber,
			"email":         ownerDAO.Email,
			"password_hash": ownerDAO.PasswordHash,
			"updated_at":    time.Now(),
		}).Error
}

func (bo *BusinessOwnerRepository) Upsert(conn abstract.IDBConnection, owner *domain.BusinessOwner) error {
	db := conn.Get().(*gorm.DB)

	ownerDAO := &model.BusinessOwner{}
	ownerDAO, err := ownerDAO.ToModel(owner)
	if err != nil {
		return err
	}

	var existing model.BusinessOwner
	err = db.Where("inn = ? OR phone_number = ? OR email = ?",
		ownerDAO.INN, ownerDAO.PhoneNumber, ownerDAO.Email).First(&existing).Error

	if err == nil {
		ownerDAO.ID = existing.ID
		return db.Save(ownerDAO).Error
	}

	return db.Create(ownerDAO).Error
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
		Select("password_hash").
		Scan(&hashedPassword).Error
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

func (bo *BusinessOwnerRepository) GetByLogin(conn abstract.IDBConnection, login string) (*domain.BusinessOwner, error) {
	db := conn.Get().(*gorm.DB)

	var ownerDAO model.BusinessOwner
	err := db.Where("email = ? OR phone_number = ?", login, login).First(&ownerDAO).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return ownerDAO.ToDomain()
}
