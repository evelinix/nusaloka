package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type Referal struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	Code string	`gorm:"uniqueIndex:idx_referals_code"`
	UserID uuid.UUID 
	User User `gorm:"foreignKey:UserID"`
}

func (Referal) TableName() string {
	return "account_referals"
}

func (r *Referal) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New()
	return
}