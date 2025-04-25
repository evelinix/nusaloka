package dto

type StartLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type StartLoginResponse struct {
	Challenge string `json:"challenge"`
}

type FinishLoginRequest struct {
	CredentialID string `json:"credential_id" validate:"required"`
	ClientData   string `json:"client_data" validate:"required"`
	AuthenticatorData string `json:"authenticator_data" validate:"required"`
	Signature string `json:"signature" validate:"required"`
}

type FinishRegistrationRequest struct {
	CredentialID string `json:"credential_id" validate:"required"`
	ClientData   string `json:"client_data" validate:"required"`
	AuthenticatorData string `json:"authenticator_data" validate:"required"`
	Signature string `json:"signature" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type StartRegistrationRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type StartRegistrationResponse struct {
	Challenge string `json:"challenge"`
}