package response

import (
	"barcelonaZoo/pkg/model"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

func SerializeUser(u *model.User) *User {
	var icon string
	if u.Icon.Valid == true {
		icon = u.Icon.String
	} else {
		icon = "default icon url" // 書き換える
	}

	return &User{
		Id:   u.ID,
		Name: u.Name,
		Icon: icon,
	}
}
