package response

import (
	"time"
)

type ChallengeRecord struct {
	Id             int    `json:"id"`
	Content        string `json:"content"`
	Comment        string `json:"comment"`
	ChallengeTheme struct {
		Id          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		User        struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
			Icon string `json:"icon"`
		} `json:"user"`
		Recordable  bool      `json:"recordable"`
		IsInt       bool      `json:"is_int"`
		Unit        bool      `json:"unit"`
		RankingType string    `json:"rankint_type"`
		CreatedAt   time.Time `json:"created_at"`
	} `json:"challenge_theme"`
	User struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Icon string `json:"icon"`
	} `json:"user"`
	Record    float32   `json:"record"`
	CreatedAt time.Time `json:"created_at"`
}
