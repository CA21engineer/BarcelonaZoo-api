package user

import (
	"barcelonaZoo/api/middleware"
	"barcelonaZoo/api/requestBody"
	"barcelonaZoo/pkg/service/user"
	"errors"
	"net/http"
	"strconv"

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

func GetUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "user id must be int"})
		return
	}

	if u, err := user.GetUser(ctx, id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "not found"})
	} else {
		ctx.JSON(http.StatusOK, u)
	}
}
