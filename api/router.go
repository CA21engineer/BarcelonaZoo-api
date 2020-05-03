package api

import (
	"barcelonaZoo/api/middleware"

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

	// firebase middlewareの作成
	firebaseClient := middleware.CreateFirebaseInstance()

	// test api
	testGroup := r.Group("/test")
	testRouter(testGroup)

	// user api
	userGroup := r.Group("/users")
	userGroup.Use(firebaseClient.MiddlewareFunc())
	userRouter(userGroup)

	return r
}
