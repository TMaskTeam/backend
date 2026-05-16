package model

import (
	"backend/src/internal/domain"
	"time"
)

type BonusProgram struct {
	ProgramID   int       `gorm:"column:program_id;primaryKey"`
	BusinessID  int       `gorm:"column:business_id;foreignKey"`
	ProgramName string    `gorm:"column:program_name;not null"`
	TokenName   string    `gorm:"column:token_name;not null"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (m *BonusProgram) ToDomain() (*domain.BonusProgram, error) {
	return ToDomain[BonusProgram, domain.BonusProgram](m)
}

func (m *BonusProgram) ToModel(domainObj *domain.BonusProgram) (*BonusProgram, error) {
	return ToModel[BonusProgram, domain.BonusProgram](domainObj)
}
