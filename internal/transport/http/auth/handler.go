package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	STATUS_OK    = "ok"
	STATUS_ERROR = "error"
)

type Handler struct {
	authUsecase AuthUsecaseInterface
}

func NewAuthHandler(authUseCase AuthUsecaseInterface) *Handler {
	return &Handler{
		authUsecase: authUseCase,
	}
}

type SignInInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type SignUpInput struct {
	Login     string `json:"login"`
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
	Password  string `json:"password"`
}

type response struct {
	Status string `json:"status"`
	Msg    string `json:"message,omitempty"`
}

func newResponse(status, msg string) *response {
	return &response{
		Status: status,
		Msg:    msg,
	}
}

type signInResponse struct {
	*response
	Token string `json:"token,omitempty"`
}

func newSignInResponse(status, msg, token string) *signInResponse {
	return &signInResponse{
		&response{
			Status: status,
			Msg:    msg,
		},
		token,
	}
}

/*Регистрация*/
func (a *Handler) SignUp(c *gin.Context) {
	inp := new(SignUpInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	_, err := a.authUsecase.SingUp(inp.Login, inp.FirstName, inp.LastName, inp.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, newResponse(STATUS_ERROR, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

/*Авторизация*/
func (a *Handler) SignIn(c *gin.Context) {
	inp := new(SignInInput)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := a.authUsecase.SingIn(inp.Login, inp.Password)

	if err != nil {
		handleTokenError(c, err)
		return
	}

	c.JSON(http.StatusOK, newSignInResponse(STATUS_OK, "", token))

}
