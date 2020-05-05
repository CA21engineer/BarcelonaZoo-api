package api

import (
	"barcelonaZoo/pkg/controller/challengetheme"

	"github.com/gin-gonic/gin"
)

func createChallengeThemeRouter(r *gin.RouterGroup) {
	r.POST("", challengetheme.CreateChallengeTheme)
}
