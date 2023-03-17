package auth

import (
	"github.com/gin-gonic/gin"
	"pedigree/internal/usecase"
)

func RegisterHTTPEndpoints(router *gin.Engine, ac usecase.AuthUsecase) {
	h := NewAuthHandler(&ac)
	authEndpoints := router.Group("/auth")
	{
		authEndpoints.POST("sign-up", h.SignUp)
		authEndpoints.POST("sign-in", h.SignIn)
	}
}
