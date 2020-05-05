package requestBody

type CreateChallengeTheme struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Recordable  bool   `json:"recordable"`
	IsInt       bool   `json:"is_int"`
	Unit        string `json:"unit"`
	RankingType int    `json:"ranking_type"`
}
