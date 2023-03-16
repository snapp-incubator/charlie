package model

// Submit model.
// each question can have one or more submits.
// each submit is for a user.
// each submit has a file that is the combination of ID+UserID.
type Submit struct {
	Base
	UserID    uint   `json:"user_id"`
	Score     int    `json:"score"`
	Extension string `json:"extension"`
}
