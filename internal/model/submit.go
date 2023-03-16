package model

import (
	"github.com/amirhnajafiz/DJaaS/pkg/enum"
)

// Submit model.
// each question can have one or more submits.
// each submit is for a user.
// each submit has a file that is the combination of ID+UserID.
type Submit struct {
	Base
	QuestionID uint        `json:"question_id"`
	UserID     uint        `json:"user_id"`
	Score      int         `json:"score"`
	Extension  string      `json:"extension"`
	Status     enum.Status `json:"status"`
}
