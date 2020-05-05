package user

import (
	"barcelonaZoo/api/response"
	"barcelonaZoo/db"
	"barcelonaZoo/pkg/model"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func CreateNewUser(ctx context.Context, uid, name string) error {
	newUser := &model.User{
		UID:  uid,
		Name: name,
	}

	if err := newUser.Insert(ctx, db.DB, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func GetUser(ctx *gin.Context, id int) (*response.User, error) {
	u, err := model.FindUser(ctx, db.DB, id)
	if err != nil {
		return &response.User{}, err
	}

	return response.SerializeUser(u), nil
}
