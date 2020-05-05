package challengetheme

import (
	"barcelonaZoo/api/middleware"
	"barcelonaZoo/api/requestBody"
	"barcelonaZoo/pkg/service/challengetheme"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateChallengeTheme(ctx *gin.Context) {
	uid, ok := ctx.Get(middleware.AuthCtxKey)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("認証されていません。再度ログインしてください。").Error())
		return
	}

	// 画像/動画を取得
	content, fileHeader, err := ctx.Request.FormFile("content")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}

	// jsonを取得
	var reqBody requestBody.CreateChallengeTheme
	jsonStr := ctx.Request.FormValue("challenge_theme")
	if err := json.Unmarshal([]byte(jsonStr), &reqBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	newChallengeData, err := challengetheme.CreateNewChallengeTheme(
		ctx,
		uid.(string),
		reqBody.Title,
		reqBody.Description,
		reqBody.Unit,
		reqBody.Recordable,
		reqBody.IsInt,
		reqBody.RankingType,
		content,
		fileHeader,
	)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, newChallengeData)
	return
}
