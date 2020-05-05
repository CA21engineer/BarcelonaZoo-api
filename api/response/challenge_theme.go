package response

import (
	"barcelonaZoo/pkg/model"

	"github.com/volatiletech/null"
)

type ChallengeTheme struct {
	Id          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	User        User        `json:"user"`
	Recordable  bool        `json:"recordable"`
	IsInt       null.Bool   `json:"is_int"`
	Unit        null.String `json:"unit"`
	RankingType null.Int8   `json:"rankint_type"`
	CreatedAt   string      `json:"created_at"`
}

func SerializeChallengeTheme(t *model.ChallengeTheme, u *model.User) *ChallengeTheme {
	return &ChallengeTheme{
		Id:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		User:        *SerializeUser(u),
		Recordable:  t.Recordable,
		IsInt:       t.IsInt,
		Unit:        t.Unit,
		RankingType: t.RankingType,
		CreatedAt:   t.CreatedAt.Format("2000-01-01 00:00:00"),
	}
}
