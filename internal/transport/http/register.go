package http

import (
	"github.com/gin-gonic/gin"
	"pedigree/internal/usecase"
)

func RegisterHTTPEndpoints(router *gin.Engine, ac usecase.AuthUseCase) {

	//Эндпоинты авторизации
	h := NewAuthHandler(&ac)
	authEndpoints := router.Group("/auth")
	{
		authEndpoints.POST("sign-up", h.SignUp)
		//	authEndpoints.POST("sign-in", h.SignUp)
	}
}
