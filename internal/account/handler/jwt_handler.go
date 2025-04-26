package handler

import (
	"net/http"

	"github.com/evelinix/nusaloka/internal/shared/jwtutil"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/rs/zerolog/log"
)

func JWKSHandler() gin.HandlerFunc {
	pub := jwtutil.GetPublicKey()

	jwkKey, err := jwk.New(pub)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Failed to create JWK key")
	}
	if err := jwkKey.Set(jwk.KeyIDKey, "nusaloka"); err != nil {
		log.Error().Stack().Err(err).Msg("Failed to set JWK key ID")
	}
	if err := jwkKey.Set(jwk.AlgorithmKey, "ES512"); err != nil {
		log.Error().Stack().Err(err).Msg("Failed to set JWK algorithm")
	}
	if err := jwkKey.Set(jwk.KeyUsageKey, "sig"); err != nil {
		log.Error().Stack().Err(err).Msg("Failed to set JWK key usage")
	}

	set := jwk.NewSet()
	set.Add(jwkKey)

	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, set)
	}
}
