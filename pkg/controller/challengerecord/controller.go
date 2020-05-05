package challengerecord

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateChallengeRecord(ctx *gin.Context) {

}

func GetChallengeRecords(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "challenge theme id must be int"})
		return
	}
}
