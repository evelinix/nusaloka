package helper

import (
	"encoding/binary"
	"errors"
)

type AuthenticatorData struct {
	RPIDHash     []byte
	Flags        byte
	SignCount    uint32
	AAGUID       []byte
	CredentialID []byte
	PublicKey    []byte
}

func ParseAuthenticatorData(authData []byte) (*AuthenticatorData, error) {
	if len(authData) < 37 {
		return nil, errors.New("authenticatorData too short")
	}

	rpIDHash := authData[0:32]
	flags := authData[32]
	signCount := binary.BigEndian.Uint32(authData[33:37])

	rest := authData[37:]

	if flags&0x40 == 0 { // 0x40 = AT (Attested Credential Data) flag
		return &AuthenticatorData{
			RPIDHash:  rpIDHash,
			Flags:     flags,
			SignCount: signCount,
		}, nil
	}

	if len(rest) < 18 {
		return nil, errors.New("attested credential data too short")
	}

	aaguid := rest[0:16]
	credIDLen := binary.BigEndian.Uint16(rest[16:18])

	if len(rest) < int(18+credIDLen) {
		return nil, errors.New("credential ID length overflow")
	}

	credID := rest[18 : 18+credIDLen]
	publicKey := rest[18+credIDLen:]

	return &AuthenticatorData{
		RPIDHash:     rpIDHash,
		Flags:        flags,
		SignCount:    signCount,
		AAGUID:       aaguid,
		CredentialID: credID,
		PublicKey:    publicKey,
	}, nil
}