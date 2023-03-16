package repository

import (
	"github.com/amirhnajafiz/DJaaS/internal/model"
)

type UserRepository interface {
	Create(user *model.User) error
	Get(username string) (user *model.User, err error)
	GetAll() ([]*model.User, error)
	Delete(username string) error
	Update(user *model.User) error
}
