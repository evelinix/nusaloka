package config

import (
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/rs/zerolog/log"
)

var (
	webAuthn *webauthn.WebAuthn
	err error
)

func InitWebAuthn() {
	wconfig := &webauthn.Config{
		RPDisplayName: "Nusaloka", // required
		RPID:          "localhost", //
		RPOrigins:     []string{"http://localhost:9001"}, // frontend
	}

	if webAuthn, err = webauthn.New(wconfig); err != nil {
		log.Err(err).Msg("Failed to initialize webauthn")
	}
}

func GetWebAuthn() *webauthn.WebAuthn {
	return webAuthn
}