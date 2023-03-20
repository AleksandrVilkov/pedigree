package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, ac AuthUsecaseInterface) {

	h := NewAuthHandler(&ac)
	{
		router.POST("/signUp", h.SignUp)
		router.POST("/signIn", h.SignIn)
	}
}
