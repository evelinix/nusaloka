package service

import (
	"context"

	"github.com/evelinix/nusaloka/internal/account/dto"
)

type AuthServiceInterface interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type AccountServiceInterface interface {
	UpdatePassword(ctx context.Context, req dto.UpdatePasswordRequest) error //update password
	UpdateAvatar(ctx context.Context, req dto.UpdateAvatarRequest) error //update avatar
	GetAccount(ctx context.Context) (dto.AccountResponse, error) //get account
}

type ReferalServiceInterface interface {
	GetReferal(ctx context.Context) (dto.ReferalListResponse, error) //get referal list
	GetReferalCode(ctx context.Context) (dto.ReferalCodeResponse, error) //get referal code
}

type WebAuthnServiceInterface interface {
	WebAuthnStartLogin(ctx context.Context, req dto.BeginLoginRequest) (dto.BeginLoginRequest, error) //Start WebAuth Login
	WebAuthnFinishLogin(ctx context.Context, req dto.FinishLoginRequest) error //Finish WebAuth Login
	WebAuthnStartRegistration(ctx context.Context, req dto.BeginRegistrationRequest) (dto.BeginRegistrationResponse, error) //Start WebAuth Register
	WebAuthnFinishRegistration(ctx context.Context, req dto.FinishRegistrationRequest) error //Finish WebAuth Register
}