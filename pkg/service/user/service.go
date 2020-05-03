package user

import (
	"barcelonaZoo/db"
	"barcelonaZoo/pkg/model"
	"context"

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
