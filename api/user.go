package api

import (
	"barcelonaZoo/pkg/controller/user"

	"github.com/gin-gonic/gin"
)

func userRouter(r *gin.RouterGroup) {
	r.POST("", user.CreateUser)
	r.GET("/:id", user.GetUser)
}
