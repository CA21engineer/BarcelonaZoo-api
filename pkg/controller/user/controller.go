package user

import (
	"barcelonaZoo/api/requestBody"
	"barcelonaZoo/pkg/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var reqBody requestBody.CreateUser
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := user.CreateNewUser(ctx, reqBody.Name);err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
	return
}
