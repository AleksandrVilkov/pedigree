package familyMember

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"pedigree/pkg/jwt"
	"strings"
)

const BEARER = "Bearer"

type Handler struct {
	familyMemberUsecase FmUsecaseInterface
}

func NewAuthHandler(fm FmUsecaseInterface) *Handler {
	return &Handler{
		familyMemberUsecase: fm,
	}
}

func (a *Handler) Save(c *gin.Context) {
	if checkAuth(c) != nil {
		return
	}
}

func (a *Handler) Delete(c *gin.Context) {
	if checkAuth(c) != nil {
		return
	}
}

func (a *Handler) Update(c *gin.Context) {
	if checkAuth(c) != nil {
		return
	}
}

func (a *Handler) Find(c *gin.Context) {
	if checkAuth(c) != nil {
		return
	}
}

func checkAuth(c *gin.Context) error {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return ErrEmptyToken
	}

	headerParts := strings.Split(authHeader, " ")

	if len(headerParts) != 2 || headerParts[0] != BEARER {
		c.AbortWithStatus(http.StatusUnauthorized)
		return ErrNotValidToken
	}

	err := jwt.ParseToken(headerParts[1], []byte(viper.GetString("auth.signing_key")))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return err
	}
	return nil
}
