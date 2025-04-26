package helper

import (
	"bytes"
	"fmt"

	"github.com/fxamacker/cbor/v2"
)

type AttestationObject struct {
	AuthData []byte                 `cbor:"authData"`
	Fmt      string                 `cbor:"fmt"`
	AttStmt  map[string]interface{} `cbor:"attStmt"`
}

func ParseAttestationObject(attestationBytes []byte) (*AttestationObject, error) {
	var attObj AttestationObject
	decoder := cbor.NewDecoder(bytes.NewReader(attestationBytes))
	if err := decoder.Decode(&attObj); err != nil {
		return nil, fmt.Errorf("failed to decode attestation object: %w", err)
	}
	return &attObj, nil
}