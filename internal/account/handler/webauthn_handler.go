package handler

import (
	"net/http"

	"github.com/evelinix/nusaloka/internal/account/dto"
	"github.com/evelinix/nusaloka/internal/account/service"
	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
)

type WebAuthnHandler struct {
	Service *service.WebAuthnService
}

func NewWebAuthnHandler(service *service.WebAuthnService) *WebAuthnHandler {
	return &WebAuthnHandler{
		Service: service,
	}
}

func (h *WebAuthnHandler) BeginRegister(c *gin.Context) {
	var req dto.BeginRegistrationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Service.BeginRegistration(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *WebAuthnHandler) FinishRegister(c *gin.Context) {
	var req dto.FinishRegistrationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Service.FinishRegistration(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.FinishRegistrationResponse{Status: "success"})
}

// LoginOption is used to provide parameters that modify the default Credential Assertion Payload that is sent to the user.
type LoginOption func(*protocol.PublicKeyCredentialRequestOptions)

// DiscoverableUserHandler returns a *User given the provided userHandle.
type DiscoverableUserHandler func(rawID, userHandle []byte) (user webauthn.User, err error)

func (h *WebAuthnHandler) BeginLogin(c *gin.Context) {
	var req dto.BeginLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Service.BeginLogin(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *WebAuthnHandler) FinishLogin(c *gin.Context) {
	var req dto.FinishLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Service.FinishLogin(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}