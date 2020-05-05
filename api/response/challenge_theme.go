package response

import (
	"barcelonaZoo/pkg/model"
)

type ChallengeTheme struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Description string `json:"description"`
	User        *User  `json:"user"`
	Recordable  bool   `json:"recordable"`
	IsInt       bool   `json:"is_int"`
	Unit        string `json:"unit"`
	RankingType int    `json:"rankint_type"`
	CreatedAt   string `json:"created_at"`
}

type ChallengeThemeSlice []*ChallengeTheme

func ConvertToChallengeThemeResponse(challengeThemeData *model.ChallengeTheme) *ChallengeTheme {
	if challengeThemeData.R == nil {
		return &ChallengeTheme{
			ID:          challengeThemeData.ID,
			Title:       challengeThemeData.Title,
			Content:     challengeThemeData.Content,
			Description: challengeThemeData.Description,
			User:        nil,
			Recordable:  challengeThemeData.Recordable,
			IsInt:       challengeThemeData.IsInt.Bool,
			Unit:        challengeThemeData.Unit.String,
			RankingType: int(challengeThemeData.RankingType.Int8),
			CreatedAt:   challengeThemeData.CreatedAt.Format("2006-01-02"),
		}
	}

	return &ChallengeTheme{
		ID:          challengeThemeData.ID,
		Title:       challengeThemeData.Title,
		Content:     challengeThemeData.Content,
		Description: challengeThemeData.Description,
		User:        ConvertToUserResponse(challengeThemeData.R.User),
		Recordable:  challengeThemeData.Recordable,
		IsInt:       challengeThemeData.IsInt.Bool,
		Unit:        challengeThemeData.Unit.String,
		RankingType: int(challengeThemeData.RankingType.Int8),
		CreatedAt:   challengeThemeData.CreatedAt.Format("2006-01-02"),
	}
}

func ConvertToChallengeThemeSliceResponse(theme model.ChallengeThemeSlice) ChallengeThemeSlice {
	if len(theme) == 0 {
		return nil
	}

	challengeThemeSlice := make(ChallengeThemeSlice, 0, len(theme))
	for _, challengeThemeData := range theme {
		challengeThemeSlice = append(challengeThemeSlice, ConvertToChallengeThemeResponse(challengeThemeData))
	}

	return challengeThemeSlice
}
