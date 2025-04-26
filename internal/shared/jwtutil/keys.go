package jwtutil

import (
	"crypto/ecdsa"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
)

func InitKeys() error {
	privPem, err := os.ReadFile("./keys/es512-private.pem")
	if err != nil {
		return err
	}
	pubPem, err := os.ReadFile("./keys/es512-public.pem")
	if err != nil {
		return err
	}
	privateKey, err = jwt.ParseECPrivateKeyFromPEM(privPem)
	if err != nil {
		return err
	}
	publicKey, err = jwt.ParseECPublicKeyFromPEM(pubPem)
	return err
}

func GenerateToken(sub string) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   sub,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES512, claims)
	return token.SignedString(privateKey)
}

func VerifyToken(tokenStr string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrTokenInvalidClaims
}

func GetPublicKey() *ecdsa.PublicKey {
	return publicKey
}
