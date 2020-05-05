package api

import (
	"barcelonaZoo/api/middleware"
	"barcelonaZoo/pkg/controller/challengetheme"
	"barcelonaZoo/pkg/controller/user"

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

	firebaseGroup := r.Group("")
	firebaseGroup.Use(firebaseClient.MiddlewareFunc())
	{
		// user api
		firebaseGroup.POST("/users", user.CreateUser)

		// challengetheme api
		firebaseGroup.POST("/challengetheme", challengetheme.CreateChallengeTheme)
	}

	// user api
	r.GET("/users/:id", user.GetUser)

	// challengetheme api
	r.GET("/challengetheme", challengetheme.GetChallengeThemes)

	return r
}
