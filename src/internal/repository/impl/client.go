package impl

import (
	"backend/src/internal/db/abstract"
	"backend/src/internal/domain"
	"backend/src/internal/model"

	"gorm.io/gorm"
)

type ClientRepository struct{}

func NewClientRepository() *ClientRepository {
	return &ClientRepository{}
}

func (c *ClientRepository) Upsert(conn abstract.IDBConnection, client *domain.Client) error {
	db := conn.Get().(*gorm.DB)

	clientDAO := &model.Client{}
	clientDAO, err := clientDAO.ToModel(client)
	if err != nil {
		return err
	}

	var existing model.Client
	err = db.Where("phone_number = ? OR email = ?",
		clientDAO.PhoneNumber, clientDAO.Email).First(&existing).Error

	if err == nil {
		clientDAO.ID = existing.ID
		return db.Save(clientDAO).Error
	}

	return db.Create(clientDAO).Error
}

func (c *ClientRepository) Delete(conn abstract.IDBConnection, clientID int) error {
	db := conn.Get().(*gorm.DB)
	return db.Where("client_id = ?", clientID).Delete(&model.Client{}).Error
}

func (c *ClientRepository) GetPasswordHashById(conn abstract.IDBConnection, clientID int) (string, error) {
	db := conn.Get().(*gorm.DB)

	var hashedPassword string
	err := db.Model(&model.Client{}).
		Where("client_id = ?", clientID).
		Select("password_hash").
		Scan(&hashedPassword).Error

	if err != nil {
		return "", err
	}

	return hashedPassword, err
}

func (c *ClientRepository) GetByPhoneNumber(conn abstract.IDBConnection, phoneNumber string) (*domain.Client, error) {
	db := conn.Get().(*gorm.DB)

	var clientDAO model.Client
	err := db.Where("phone_number = ?", phoneNumber).First(&clientDAO).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return clientDAO.ToDomain()
}

func (c *ClientRepository) GetByEmail(conn abstract.IDBConnection, email string) (*domain.Client, error) {
	db := conn.Get().(*gorm.DB)

	var clientDAO model.Client
	err := db.Where("email = ?", email).First(&clientDAO).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return clientDAO.ToDomain()
}

func (c *ClientRepository) GetByLogin(conn abstract.IDBConnection, login string) (*domain.Client, error) {
	db := conn.Get().(*gorm.DB)

	var clientDAO model.Client
	err := db.Where("email = ? OR phone_number = ?", login, login).First(&clientDAO).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return clientDAO.ToDomain()
}

func (c *ClientRepository) GetByID(conn abstract.IDBConnection, id int) (*domain.Client, error) {
	db := conn.Get().(*gorm.DB)

	var clientDAO model.Client
	err := db.Where("client_id = ?", id).First(&clientDAO).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return clientDAO.ToDomain()
}
