package auth

import (
	"github.com/gin-gonic/gin"
	"pedigree/internal/usecase"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, ac usecase.AuthUsecase) {

	h := NewAuthHandler(&ac)
	{
		router.POST("/signUp", h.SignUp)
		router.POST("/signIn", h.SignIn)
	}
}
