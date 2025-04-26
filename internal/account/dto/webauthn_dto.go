package dto

import "github.com/google/uuid"

// 1. Client -> Server
type BeginRegistrationRequest struct {
	UserID uuid.UUID `json:"user_id" binding:"required"`
}

// 2. Server -> Client
type BeginRegistrationResponse struct {
	Challenge          string   `json:"challenge"`
	RelyingPartyID     string   `json:"rp_id"`
	RelyingPartyName   string   `json:"rp_name"`
	UserID             string   `json:"user_id"`
	UserName           string   `json:"user_name"`
	Attestation        string   `json:"attestation"` // e.g., "direct", "indirect", "none"
	AuthenticatorSelection *AuthenticatorSelection `json:"authenticator_selection"`
	Timeout            uint64   `json:"timeout"` // optional timeout
}

// Nested DTO
type AuthenticatorSelection struct {
	AuthenticatorAttachment string `json:"authenticator_attachment"` // e.g., "platform", "cross-platform"
	RequireResidentKey      bool   `json:"require_resident_key"`
	UserVerification        string `json:"user_verification"` // e.g., "required", "preferred", "discouraged"
}

// 3. Client -> Server
type FinishRegistrationRequest struct {
	UserID        uuid.UUID `json:"user_id" binding:"required"`
	ClientDataJSON string   `json:"client_data_json" binding:"required"`
	AttestationObject string `json:"attestation_object" binding:"required"`
}

// 4. Server -> Client
type FinishRegistrationResponse struct {
	Status string `json:"status"` // success/fail
}

// 5. Client -> Server
type BeginLoginRequest struct {
	UserID uuid.UUID `json:"user_id" binding:"required"`
}

// 6. Server -> Client
type BeginLoginResponse struct {
	Challenge      string   `json:"challenge"`
	RelyingPartyID string   `json:"rp_id"`
	AllowCredentials []AllowedCredential `json:"allow_credentials,omitempty"`
	UserVerification string  `json:"user_verification"`
	Timeout         uint64   `json:"timeout"`
}

// Nested DTO
type AllowedCredential struct {
	Type string `json:"type"` // "public-key"
	ID   string `json:"id"`   // credentialID
	Transports []string `json:"transports,omitempty"` // ["usb", "nfc", "ble", "internal"]
}

// 7. Client -> Server
type FinishLoginRequest struct {
	UserID         uuid.UUID `json:"user_id" binding:"required"`
	ClientDataJSON string    `json:"client_data_json" binding:"required"`
	AuthenticatorData string `json:"authenticator_data" binding:"required"`
	Signature       string   `json:"signature" binding:"required"`
	CredentialID    string   `json:"credential_id" binding:"required"`
}

// 8. Server -> Client
type FinishLoginResponse struct {
	Status string `json:"status"` // success/fail
	Token  string `json:"token,omitempty"` // JWT or session if login success
}