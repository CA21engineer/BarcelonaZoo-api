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
		ctx.AbortWithError(http.StatusUnauthorized, errors.New("認証されていません。再度ログインしてください。"))
		return
	}

	var reqBody requestBody.CreateUser
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := user.CreateNewUser(ctx, uid.(string), reqBody.Name); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
	return
}
