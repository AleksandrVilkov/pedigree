package familyMember

import "github.com/gin-gonic/gin"

type Handler struct {
	familyMemberUsecase FmUsecaseInterface
}

func NewAuthHandler(fm FmUsecaseInterface) *Handler {
	return &Handler{
		familyMemberUsecase: fm,
	}
}

func (a *Handler) Save(c *gin.Context) {
}

func (a *Handler) Delete(c *gin.Context) {

}

func (a *Handler) Update(c *gin.Context) {

}

func (a *Handler) Find(c *gin.Context) {

}
