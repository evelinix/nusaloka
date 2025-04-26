package router

import (
	"fmt"
	"time"

	"github.com/evelinix/nusaloka/internal/account/handler"
	"github.com/evelinix/nusaloka/internal/account/observability"
	"github.com/evelinix/nusaloka/internal/account/repository"
	"github.com/evelinix/nusaloka/internal/account/service"
	"github.com/evelinix/nusaloka/internal/shared/jwtutil"
	middlewareShared "github.com/evelinix/nusaloka/internal/shared/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAccountRouter(r *gin.Engine, db *gorm.DB) {
	repo := repository.NewAuthRepository(db)
	svc := service.NewAuthService(repo)
	h := handler.NewAuthHandler(svc)

	r.GET("/health", func(c *gin.Context) {
		c.String(200, fmt.Sprintf("Account Service is up %d", time.Now().Unix()))
	})

	r.GET("/.well-known/jwks.json", handler.JWKSHandler())

	r.POST("/auth/login", h.AuthLoginHandler)

	r.POST("/auth/register", h.RegisterHandler)

	// r.POST("/auth/forgot-password", handler.ForgotPasswordHandler)

	// r.GET("/account", handler.AccountHandler)

	// r.POST("/account/update-password", handler.UpdatePasswordHandler)

	// r.POST("/account/update-avatar", handler.UpdateAvatarHandler)

	// r.GET("/referal", handler.GetReferalHandler)

	r.GET("/token", func(c *gin.Context) {
		token, _ := jwtutil.GenerateToken("test-id-123")
		c.JSON(200, gin.H{"token": token})
	})

	r.GET("/private", middlewareShared.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, authorized user!"})
	})

	r.GET("/metrics", observability.PrometheusHandler())

}
