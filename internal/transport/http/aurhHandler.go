package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pedigree/internal/transport"
)

type AuthHandler struct {
	authUseCase transport.AuthUseCaseInterface
}

func NewAuthHandler(authUseCase transport.AuthUseCaseInterface) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase,
	}
}

type SignInput struct {
	Login     string `json:"login"`
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
	Password  string `json:"password"`
}

func (a *AuthHandler) SignUp(c *gin.Context) {
	inp := new(SignInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	_, err := a.authUseCase.SingUp(c.Request.Context(), inp.Login, inp.FirstName, inp.LastName, inp.Password)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
