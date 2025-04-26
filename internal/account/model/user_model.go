package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents a user in the system.
type User struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey"`
	DisplayName string     `json:"display_name"`
	Avatar      string     `json:"avatar"`
	Username    string     `gorm:"Index:idx_users_username"`
	Email       string     `gorm:"uniqueIndex:idx_users_email;not null"`
	Password    string     `json:"-"`
	Credentials []Webauth `gorm:"foreignKey:UserID"`
	gorm.Model
}

func (User) TableName() string {
	return "account_users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
