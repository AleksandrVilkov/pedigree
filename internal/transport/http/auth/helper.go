package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleTokenError(c *gin.Context, err error) {
	if err == ErrInvalidAccessToken {
		c.AbortWithStatusJSON(http.StatusBadRequest, newSignInResponse(STATUS_ERROR, err.Error(), ""))
		return
	}

	if err == ErrUserDoesNotExist {
		c.AbortWithStatusJSON(http.StatusBadRequest, newSignInResponse(STATUS_ERROR, err.Error(), ""))
		return
	}

	c.AbortWithStatusJSON(http.StatusInternalServerError, newSignInResponse(STATUS_ERROR, err.Error(), ""))
}
