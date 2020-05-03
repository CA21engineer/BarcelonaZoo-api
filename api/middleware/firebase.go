package middleware

import (
	"barcelonaZoo/api/requestHeader"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

type FirebaseAuth interface {
	MiddlewareFunc() gin.HandlerFunc
}

type firebaseAuth struct {
	client *auth.Client
}

const (
	AuthCtxKey = "AUTHED_UID"
)

func CreateFirebaseInstance() FirebaseAuth {
	ctx := context.Background()

	// get credential of firebase
	var opt option.ClientOption
	opt = option.WithCredentialsFile("barcelona-zoo-firebase-adminsdk-as7uq-82b8457294.json")

	// firebase appの作成
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Panic(fmt.Errorf("error initializing firebase app: %v", err))
	}

	// firebase admin clientの作成
	client, err := app.Auth(ctx)
	if err != nil {
		log.Panic(fmt.Errorf("error initialize firebase instance. %v", err))
	}

	return &firebaseAuth{
		client: client,
	}
}

func (fa *firebaseAuth) MiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		fa.middlewareImpl(c)
	}
}

func (fa *firebaseAuth) middlewareImpl(c *gin.Context) {
	// Authorizationヘッダーからjwtトークンを取得
	var reqHeader requestHeader.Authorization
	if err := c.BindHeader(&reqHeader); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	jwtToken := strings.Replace(reqHeader.Token, "Bearer ", "", 1)

	// JWT の検証
	authedUserToken, err := fa.client.VerifyIDToken(c, jwtToken)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// contextにuidを格納
	c.Set(AuthCtxKey, authedUserToken.UID)
}
