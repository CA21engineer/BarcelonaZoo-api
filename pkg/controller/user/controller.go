package user

import (
	"barcelonaZoo/api/middleware"
	"barcelonaZoo/api/requestBody"
	"barcelonaZoo/pkg/service/user"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	uid, ok := ctx.Get(middleware.AuthCtxKey)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("認証されていません。再度ログインしてください。").Error())
		return
	}

	var reqBody requestBody.CreateUser
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := user.CreateNewUser(ctx, uid.(string), reqBody.Name); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusNoContent)
	return
}
