package api

import (
	"barcelonaZoo/pkg/controller/test"
	"github.com/gin-gonic/gin"
)

func testRouter(r *gin.RouterGroup) {
	r.GET("/ping", test.Ping)
}
