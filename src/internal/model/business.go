package model

import (
	"backend/src/internal/domain"
	"time"
)

type Business struct {
	BusinessID int       `gorm:"column:business_id;primaryKey"`
	OwnerID    int       `gorm:"column:owner_id;not null"`
	Name       string    `gorm:"column:name;not null"`
	Address    string    `gorm:"column:address;not null"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (m *Business) ToDomain() (*domain.Business, error) {
	return ToDomain[Business, domain.Business](m)
}

func (m *Business) ToModel(domainObj *domain.Business) (*Business, error) {
	return ToModel[Business, domain.Business](domainObj)
}

func (m *Business) ToDomainSlice(businessDAOs []Business) ([]domain.Business, error) {
	result := make([]domain.Business, 0, len(businessDAOs))
	for _, b := range businessDAOs {
		domainBusiness, err := b.ToDomain()
		if err != nil {
			return nil, err
		}
		result = append(result, *domainBusiness)
	}
	return result, nil
}
