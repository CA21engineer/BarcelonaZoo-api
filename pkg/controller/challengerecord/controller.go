package challengerecord

import (
	"barcelonaZoo/api/middleware"
	"barcelonaZoo/api/requestBody"
	"barcelonaZoo/api/response"
	"barcelonaZoo/pkg/service/challengerecord"
	"barcelonaZoo/pkg/service/challengetheme"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateChallengeRecord(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "challenge theme id must be int"})
		return
	}

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
	var reqBody requestBody.CreateChallengeRecord
	jsonStr := ctx.Request.FormValue("record_data")
	if err := json.Unmarshal([]byte(jsonStr), &reqBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	newChallengeRecord, err := challengerecord.CreateChallengeRecord(
		ctx,
		id,
		uid.(string),
		reqBody.Comment,
		reqBody.Record,
		content,
		fileHeader,
	)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, response.ConvertToChallengeRecordResponse(newChallengeRecord))
	return
}

func GetChallengeRecords(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "challenge theme id must be int"})
		return
	}

	// チャレンジテーマデータの取得
	challengethemeData, err := challengetheme.GetChallengeThemeByPK(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	// チャレンジレコード一覧取得
	slice, err := challengerecord.GetChallengeRecords(ctx, challengethemeData.ID, int(challengethemeData.RankingType.Int8))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, response.ConvertToChallengeRecordSliceResponse(slice))
	return
}
