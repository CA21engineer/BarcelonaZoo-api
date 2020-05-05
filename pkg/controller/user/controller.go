package user

import (
	"barcelonaZoo/api/middleware"
	"barcelonaZoo/api/requestBody"
	"barcelonaZoo/api/response"
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

	insertedData, err := user.CreateNewUser(ctx, uid.(string), reqBody.Name, reqBody.Icon)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, response.ConvertToUserResponse(insertedData))
	return
}

func GetUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "user id must be int"})
		return
	}

	u, err := user.GetUser(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "not found"})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, response.ConvertToUserResponse(u))
	return
}
