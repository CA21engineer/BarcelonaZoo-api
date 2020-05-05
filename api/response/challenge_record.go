package response

import (
	"barcelonaZoo/pkg/model"
)

type ChallengeRecord struct {
	ID        int     `json:"id"`
	Content   string  `json:"content"`
	Comment   string  `json:"comment"`
	User      *User   `json:"user"`
	Record    float32 `json:"record"`
	CreatedAt string  `json:"created_at"`
}

type ChallengeRecordSlice []*ChallengeRecord

func ConvertToChallengeRecordResponse(record *model.ChallengeRecord) *ChallengeRecord {
	if record.R == nil {
		return &ChallengeRecord{
			ID:        record.ID,
			Content:   record.Content,
			Comment:   record.Comment,
			User:      nil,
			Record:    record.Record.Float32,
			CreatedAt: record.CreatedAt.Format("2006-01-02"),
		}
	}

	return &ChallengeRecord{
		ID:        record.ID,
		Content:   record.Content,
		Comment:   record.Comment,
		User:      ConvertToUserResponse(record.R.User),
		Record:    record.Record.Float32,
		CreatedAt: record.CreatedAt.Format("2006-01-02"),
	}
}

func ConvertToChallengeRecordSliceResponse(slice model.ChallengeRecordSlice) ChallengeRecordSlice {
	if len(slice) == 0 {
		return nil
	}

	challengeRecordSlice := make(ChallengeRecordSlice, 0, len(slice))
	for _, challengeRecordData := range slice {
		challengeRecordSlice = append(challengeRecordSlice, ConvertToChallengeRecordResponse(challengeRecordData))
	}

	return challengeRecordSlice
}
