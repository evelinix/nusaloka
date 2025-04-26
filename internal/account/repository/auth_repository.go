package repository

import (
	"github.com/evelinix/nusaloka/internal/account/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository interface {
	FindByEmail(email string) (*model.User, error)
	FindByID(id uuid.UUID) (*model.User, error)
	Create(user *model.User) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (ur *authRepository) FindByID(id uuid.UUID) (*model.User, error) {
	var user model.User
	err := ur.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *authRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := ur.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *authRepository) Create(user *model.User) error {
	return ur.db.Create(user).Error
}