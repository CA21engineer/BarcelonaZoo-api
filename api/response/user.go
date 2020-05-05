package response

import (
	"barcelonaZoo/pkg/model"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

func ConvertToUserResponse(userData *model.User) *User {
	return &User{
		ID:   userData.ID,
		Name: userData.Name,
		Icon: userData.Icon.String,
	}
}
