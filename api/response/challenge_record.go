package response

import (
	"barcelonaZoo/pkg/model"

	"github.com/volatiletech/null"
)

type ChallengeRecord struct {
	Id        int          `json:"id"`
	Content   string       `json:"content"`
	Comment   string       `json:"comment"`
	User      User         `json:"user"`
	Record    null.Float32 `json:"record"`
	CreatedAt string       `json:"created_at"`
}

func SerializeChallengeRecord(r *model.ChallengeRecord, u *model.User) *ChallengeRecord {
	return &ChallengeRecord{
		Id:        r.ID,
		Content:   r.Content,
		Comment:   r.Comment,
		User:      *SerializeUser(u),
		Record:    r.Record,
		CreatedAt: r.CreatedAt.Format("2000-01-01 00:00:00"),
	}
}
