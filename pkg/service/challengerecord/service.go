package challengerecord

import (
	"barcelonaZoo/db"
	"barcelonaZoo/pkg/model"
	"barcelonaZoo/pkg/service/uploader"
	"mime/multipart"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

const (
	asc  = 0
	desc = 1
)

func CreateChallengeRecord(
	ctx *gin.Context,
	id int,
	uid, comment string,
	record float32,
	content multipart.File,
	fileHeader *multipart.FileHeader,
) (*model.ChallengeRecord, error) {
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

	// 新規チャレンジレコード作成
	newChallengeRecord := &model.ChallengeRecord{
		Content:          contentURL,
		Comment:          comment,
		ChallengeThemeID: id,
		UserID:           userData.ID,
		Record:           null.Float32From(record),
		CreatedAt:        time.Now(),
	}
	if err := newChallengeRecord.Insert(ctx, db.DB, boil.Infer()); err != nil {
		return nil, err
	}

	return newChallengeRecord, nil
}

func GetChallengeRecords(ctx *gin.Context, id, rankingType int) (model.ChallengeRecordSlice, error) {
	var orderqm qm.QueryMod
	switch rankingType {
	case asc:
		orderqm = qm.OrderBy("challenge_records.record")
	case desc:
		orderqm = qm.OrderBy("challenge_records.record DESC")
	default:
		orderqm = qm.OrderBy("challenge_records.created_at")
	}

	slice, err := model.ChallengeRecords(
		qm.Where("challenge_theme_id=?", id),
		orderqm,
	).All(ctx, db.DB)
	if err != nil {
		return nil, err
	}

	return slice, nil
}
