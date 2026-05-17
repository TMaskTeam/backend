package model

import (
	"backend/src/internal/domain"
	"time"
)

type BonusProgramInfo struct {
	ProgramInfoID            int       `gorm:"column:program_info_id;primaryKey"`
	ProgramID                int       `gorm:"column:program_id;not null"`
	VisitTokens              int       `gorm:"column:visit_tokens;not null"`
	PercentagePurchaseTokens int       `gorm:"column:percentage_purchase_tokens;not null"`
	RegisterTokens           int       `gorm:"column:register_tokens;not null"`
	BirthdayTokens           int       `gorm:"column:birthday_tokens;not null"`
	FriendInviteTokens       int       `gorm:"column:friend_invite_tokens;not null"`
	MinimumReceiptLimit      int       `gorm:"column:minimum_receipt_limit;not null"`
	CreatedAt                time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt                time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (m *BonusProgramInfo) ToDomain() (*domain.BonusProgramInfo, error) {
	return ToDomain[BonusProgramInfo, domain.BonusProgramInfo](m)
}

func (m *BonusProgramInfo) ToModel(domainObj *domain.BonusProgramInfo) (*BonusProgramInfo, error) {
	return ToModel[BonusProgramInfo, domain.BonusProgramInfo](domainObj)
}
