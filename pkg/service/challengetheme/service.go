package challengetheme

import (
	"barcelonaZoo/db"
	"barcelonaZoo/pkg/model"
	"barcelonaZoo/pkg/service/uploader"
	"context"
	"mime/multipart"
	"time"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func CreateNewChallengeTheme(
	ctx context.Context,
	uid, title, description, unit string,
	recordable, isInt bool,
	rankingType int,
	content multipart.File,
	fileHeader *multipart.FileHeader,
) (*model.ChallengeTheme, error) {
	// uidを元にユーザの取得
	userData, err := model.Users(model.UserWhere.UID.EQ(uid)).One(ctx, db.DB)
	if err != nil {
		return nil, err
	}

	// コンテンツのアップロード
	contentURL, err := uploader.ContentUploadToGCS(ctx, content, fileHeader)
	if err != nil {
		return nil, err
	}

	// 新規チャレンジネタ作成
	newChallengeTheme := &model.ChallengeTheme{
		Title:       title,
		Content:     contentURL,
		Description: description,
		UserID:      userData.ID,
		Recordable:  recordable,
		IsInt:       null.BoolFrom(isInt),
		Unit:        null.StringFrom(unit),
		RankingType: null.Int8From(int8(rankingType)),
		CreatedAt:   time.Now(),
	}
	if err := newChallengeTheme.Insert(ctx, db.DB, boil.Infer()); err != nil {
		return nil, err
	}

	// トランザクション終了
	return newChallengeTheme, nil
}

func GetChallengeThemes(ctx context.Context) (model.ChallengeThemeSlice, error) {
	challengeThemeSlice, err := model.ChallengeThemes(
		qm.Load("User"),
		qm.OrderBy("challenge_themes.created_at"),
	).All(ctx, db.DB)
	if err != nil {
		return nil, err
	}
	return challengeThemeSlice, nil
}
