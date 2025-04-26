package repository

import (
	"github.com/evelinix/nusaloka/internal/account/model"
	"gorm.io/gorm"
)

type WebAuthnRepository interface {
	GetUserByID(id string) (*model.User, error)
	StoreCredential(cred model.Webauth) error
	FindAllByUserID(userID string) ([]model.Webauth, error)
}

type webAuthnRepository struct {
	DB *gorm.DB
}

func NewWebAuthnRepository(db *gorm.DB) WebAuthnRepository {
	return &webAuthnRepository{db}
}

// GetUserByID implements WebAuthnRepository.
func (w *webAuthnRepository) GetUserByID(userID string) (*model.User, error) {
	var user model.User
	err := w.DB.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// StoreCredential implements WebAuthnRepository.
func (wr *webAuthnRepository) StoreCredential(cred model.Webauth) error {
	return wr.DB.Create(&cred).Error
}

// FindAllByUserID implements WebAuthnRepository.
func (wr *webAuthnRepository) FindAllByUserID(userID string) ([]model.Webauth, error) {
	var creds []model.Webauth
	if err := wr.DB.Where("user_id = ?", userID).Find(&creds).Error; err != nil {
		return nil, err
	}
	return creds, nil
}


