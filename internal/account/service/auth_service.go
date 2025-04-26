package service

import (
	"context"
	"errors"

	"github.com/evelinix/nusaloka/internal/account/dto"
	"github.com/evelinix/nusaloka/internal/account/model"
	"github.com/evelinix/nusaloka/internal/account/repository"
	"github.com/evelinix/nusaloka/internal/shared/jwtutil"
	"github.com/evelinix/nusaloka/internal/shared/utils"
)

type AuthService struct {
	userRepo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthServiceInterface {
	return &AuthService{repo}
}

func (as *AuthService) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	_, err := as.userRepo.FindByEmail(req.Email)
	if err == nil {
		return nil, errors.New("email sudah terdaftar")
	}

	HashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email:    req.Email,
		Password: HashPassword,
	}

	err = as.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return &dto.RegisterResponse{
		ID:    user.ID.String(),
		Email: user.Email,
	}, nil
}

func (as *AuthService) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := as.userRepo.FindByEmail(req.Email)
	if err != nil || !utils.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("Your email or password is incorrect")
	}
	
	tokem, err := jwtutil.GenerateToken(user.ID.String())
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Token: tokem,
	}, nil
}