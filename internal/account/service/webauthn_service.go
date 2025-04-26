package service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/evelinix/nusaloka/internal/account/config"
	"github.com/evelinix/nusaloka/internal/account/dto"
	"github.com/evelinix/nusaloka/internal/account/helper"
	"github.com/evelinix/nusaloka/internal/account/model"
	"github.com/evelinix/nusaloka/internal/account/repository"
	"github.com/google/uuid"
)

type WebAuthnService struct {
	AuthRepo repository.AuthRepository
	CredentialRepo repository.WebAuthnRepository
	Config config.Config
}

func NewWebAuthnService(raw repository.AuthRepository, rwa repository.WebAuthnRepository, cfg config.Config) *WebAuthnService {
	return &WebAuthnService{
		AuthRepo:   raw,
		CredentialRepo: rwa,
		Config:     cfg,
	}
}

func (s *WebAuthnService) BeginRegistration(req dto.BeginRegistrationRequest) (dto.BeginRegistrationResponse, error) {
	user, err := s.AuthRepo.FindByID(req.UserID)
	if err != nil {
		return dto.BeginRegistrationResponse{}, err
	}
	
	// challenge := generateChallenge()
	challengeBytes := make([]byte, 32) // 32 bytes = 256 bit
	if _, err := rand.Read(challengeBytes); err != nil {
		return dto.BeginRegistrationResponse{}, errors.New("failed to generate secure random challenge")
	}
	challengeBase64 := base64.RawURLEncoding.EncodeToString(challengeBytes)

	if user.Username == "" {
		user.Username = "user_" + user.ID.String()
	}

	resp := dto.BeginRegistrationResponse{
		Challenge: challengeBase64,
		RelyingPartyID: s.Config.RPID,
		RelyingPartyName: s.Config.RPName,
		UserID: user.ID.String(),
		UserName: user.Username,
		Attestation: "none",
		AuthenticatorSelection: &dto.AuthenticatorSelection{
			AuthenticatorAttachment: "platform",
			RequireResidentKey: false,
			UserVerification: "preferred",
		},
		Timeout: 60000,
	}

	// TODO: Save challenge to cache/store

	return resp, nil
}

func (s *WebAuthnService) FinishRegistration(req dto.FinishRegistrationRequest) error {
	attestationBytes, err := helper.Base64URLDecode(req.AttestationObject)
	if err != nil {
		return fmt.Errorf("invalid attestation object: %w", err)
	}

	attObj, err := helper.ParseAttestationObject(attestationBytes)
	if err != nil {
		return fmt.Errorf("failed to parse attestation object: %w", err)
	}

	authData, err := helper.ParseAuthenticatorData(attObj.AuthData)
	if err != nil {
		return fmt.Errorf("failed to parse authenticator data: %w", err)
	}

	credential := model.Webauth{
		ID:              uuid.NewString(),
		UserID:          req.UserID,
		CredentialID:    authData.CredentialID,
		PublicKey:       authData.PublicKey,
		AAGUID:          helper.Base64URLEncode(authData.AAGUID),
		AttestationType: attObj.Fmt,
		SignCount:       authData.SignCount,
	}

	return s.CredentialRepo.StoreCredential(credential)
}


func (s *WebAuthnService) BeginLogin(req dto.BeginLoginRequest) (dto.BeginLoginResponse, error) {
	creds, err := s.CredentialRepo.FindAllByUserID(req.UserID.String())
	if err != nil || len(creds) == 0 {
		return dto.BeginLoginResponse{}, errors.New("no credentials found")
	}

	var allowed []dto.AllowedCredential
	for _, cred := range creds {
		allowed = append(allowed, dto.AllowedCredential{
			Type: "public-key",
			ID: base64.RawURLEncoding.EncodeToString([]byte(cred.ID)),
			Transports: []string{"internal"},
		})
	}

	return dto.BeginLoginResponse{
		Challenge: generateChallenge(),
		RelyingPartyID: s.Config.RPID,
		AllowCredentials: allowed,
		UserVerification: "preferred",
		Timeout: 60000,
	}, nil
}

func (s *WebAuthnService) FinishLogin(req dto.FinishLoginRequest) (dto.FinishLoginResponse, error) {
	// TODO: Validate clientDataJSON, authenticatorData, signature
	return dto.FinishLoginResponse{
		Status: "success",
		Token: "dummy_token",
	}, nil
}

func generateChallenge() string {
	return base64.RawURLEncoding.EncodeToString([]byte(uuid.NewString()))
}
