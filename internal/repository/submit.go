package repository

import "github.com/amirhnajafiz/DJaaS/internal/model"

type SubmitRepository interface {
	Create(submit *model.Submit) error
	GetQuestionSubmits(questionID uint) ([]*model.Submit, error)
	GetQuestionSubmitsOfUser(questionID, userID uint) ([]*model.Submit, error)
}
