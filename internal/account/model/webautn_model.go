package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Webauth struct {
	ID              string    `gorm:"primaryKey"`      // WebauthnID (base64url encoded)
	UserID          uuid.UUID `gorm:"type:uuid;index"` // Foreign key ke User
	PublicKey       []byte    `gorm:"not null"`
	CredentialID    []byte
	AttestationType string `gorm:"not null"`
	AAGUID          string `gorm:"size:36;not null"`
	SignCount       uint32 `gorm:"not null"`
	BackupEligible  bool   `gorm:"not null;default:false"`
	BackupState     bool   `gorm:"not null;default:false"`
	DevicePublicKey []byte // optional field
	AuthTransport   string `gorm:"type:varchar(255)"` // Store AuthenticatorTransport as string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	User            User `gorm:"foreignKey:UserID"` // Relasi dengan User
	gorm.Model
}

func (Webauth) TableName() string {
	return "account_credentials"
}
