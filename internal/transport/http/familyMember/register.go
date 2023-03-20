package familyMember

import (
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, fm FmUsecaseInterface) {

	h := NewAuthHandler(fm)
	{
		router.POST("/save", h.Save)
		router.GET("/find", h.Find)
		router.GET("/delete", h.Delete)
		router.POST("/update", h.Update)
	}
}
