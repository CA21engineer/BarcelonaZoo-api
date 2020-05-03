package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	// CORS対応
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	r.Use(cors.New(config))

	// test api
	testGroup := r.Group("/test")
	testRouter(testGroup)

	// user api
	userGroup := r.Group("/users")
	userRouter(userGroup)

	return r
}