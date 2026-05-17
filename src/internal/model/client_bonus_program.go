package model

import (
	"backend/src/internal/domain"
	"time"
)

type ClientBonusProgram struct {
	ID          int       `gorm:"column:client_bonus_program_id;primaryKey"`
	ProgramID   int       `gorm:"column:program_id;not null"`
	ClientID    int       `gorm:"column:client_id;not null"`
	TokensCount int       `gorm:"column:tokens;not null"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (modelObj *ClientBonusProgram) ToDomain() (*domain.ClientBonusProgram, error) {
	return ToDomain[ClientBonusProgram, domain.ClientBonusProgram](modelObj)
}

func (modelObj *ClientBonusProgram) ToModel(domainObj *domain.ClientBonusProgram) (*ClientBonusProgram, error) {
	return ToModel[ClientBonusProgram, domain.ClientBonusProgram](domainObj)
}
