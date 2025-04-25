package store

import "github.com/evelinix/nusaloka/internal/account/model"

type UserStore interface {
	CreateUser(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}