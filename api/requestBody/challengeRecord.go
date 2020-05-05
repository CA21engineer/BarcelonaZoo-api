package requestBody

type CreateChallengeRecord struct {
	Comment string  `json:"comment"`
	Record  float32 `json:"record"`
}
