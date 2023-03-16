package repository

import (
	"github.com/amirhnajafiz/DJaaS/internal/model"
)

type QuestionRepository interface {
	Create(question *model.Question) error
	Delete(questionID uint) error
	GetSingle(questionID uint) (*model.Question, error)
	GetClassQuestions(classID uint) ([]*model.Question, error)
	GetAll() ([]*model.Question, error)
}
