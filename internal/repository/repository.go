package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	Class    ClassRepository
	User     UserRepository
	Question QuestionRepository
	Submit   SubmitRepository
}

func New(db *gorm.DB) *Repository {
	return &Repository{}
}
