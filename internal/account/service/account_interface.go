package service

import (
	"context"

	"github.com/evelinix/nusaloka/internal/account/dto"
)

type AuthService interface {
	Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error )
	Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)
}

type AccountService interface {
	UpdatePassword(ctx context.Context, req dto.UpdatePasswordRequest) error //update password
	UpdateAvatar(ctx context.Context, req dto.UpdateAvatarRequest) error //update avatar
	GetAccount(ctx context.Context) (dto.AccountResponse, error) //get account
}

type ReferalService interface {
	GetReferal(ctx context.Context) (dto.ReferalListResponse, error) //get referal list
	GetReferalCode(ctx context.Context) (dto.ReferalCodeResponse, error) //get referal code
}

type WebAuthnService interface {
	StartLogin(ctx context.Context, req dto.StartLoginRequest) (dto.StartLoginResponse, error) //Start WebAuth Login
	FinishLogin(ctx context.Context, req dto.FinishLoginRequest) error //Finish WebAuth Login
	StartRegistration(ctx context.Context, req dto.StartRegistrationRequest) (dto.StartRegistrationResponse, error) //Start WebAuth Register
	FinishRegistration(ctx context.Context, req dto.FinishRegistrationRequest) error //Finish WebAuth Register
}