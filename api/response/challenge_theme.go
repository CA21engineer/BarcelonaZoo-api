package response

import (
	"time"
)

type ChallengeTheme struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	User        struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"user"`
	Recordable  bool      `json:"recordable"`
	IsInt       bool      `json:"is_int"`
	Unit        bool      `json:"unit"`
	RankingType string    `json:"rankint_type"`
	CreatedAt   time.Time `json:"created_at"`
}
