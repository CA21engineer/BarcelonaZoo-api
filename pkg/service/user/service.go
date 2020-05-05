package user

import (
	"barcelonaZoo/db"
	"barcelonaZoo/pkg/model"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
)

func CreateNewUser(ctx context.Context, uid, name, icon string) (*model.User, error) {
	newUser := &model.User{
		UID:  uid,
		Name: name,
		Icon: null.StringFrom(icon),
	}

	if err := newUser.Insert(ctx, db.DB, boil.Infer()); err != nil {
		return nil, err
	}

	return newUser, nil
}

func GetUser(ctx *gin.Context, id int) (*model.User, error) {
	u, err := model.FindUser(ctx, db.DB, id)
	if err != nil {
		return nil, err
	}

	return u, nil
}
