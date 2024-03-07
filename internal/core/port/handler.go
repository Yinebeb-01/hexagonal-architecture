package port

import (
	"github.com/gin-gonic/gin"
)

// Todo: better if interfaces are generic and protocol independent

// VideoHandler for RESTful API
type VideoHandler interface {
	Save(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	ShowAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(Ctx *gin.Context)
}