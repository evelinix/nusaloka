package router

import (
	"github.com/evelinix/nusaloka/internal/account/config"
	"github.com/evelinix/nusaloka/internal/account/handler"
	"github.com/evelinix/nusaloka/internal/account/repository"
	"github.com/evelinix/nusaloka/internal/account/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupWebAuthnRoutes(r *gin.Engine, db *gorm.DB) {
	repoUser := repository.NewAuthRepository(db)
	repoWebauth := repository.NewWebAuthnRepository(db)
	svc := service.NewWebAuthnService(repoUser , repoWebauth, *config.AccountConfig)
	h := handler.NewWebAuthnHandler(svc)
	
	group := r.Group("/webauthn")
	{
		group.GET("/webauthn.html", func(ctx *gin.Context) {
			ctx.File("./static/webauthn.html")
		})
		group.POST("/register/begin", h.BeginRegister)
		group.POST("/register/finish", h.FinishRegister)
		group.POST("/login/begin", h.BeginLogin)
		group.POST("/login/finish", h.FinishLogin)
	}
}