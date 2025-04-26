package handler

import (
	"net/http"

	"github.com/evelinix/nusaloka/internal/account/dto"
	"github.com/evelinix/nusaloka/internal/account/service"
	"github.com/evelinix/nusaloka/internal/shared/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthServiceInterface
}

func NewAuthHandler(s service.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		authService: s,
	}
}

func (ah *AuthHandler) AuthLoginHandler(c *gin.Context){
	user, ok := utils.BindAndValidate[dto.LoginRequest](c)
	if !ok {
		return
	}

	res, err := ah.authService.Login(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ah *AuthHandler) RegisterHandler(c *gin.Context) {
	user, ok := utils.BindAndValidate[dto.RegisterRequest](c)
	if !ok {
		return
	}

	res, err := ah.authService.Register(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
	
}
