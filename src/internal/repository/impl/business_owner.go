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
