package repository

import (
	"github.com/amirhnajafiz/DJaaS/internal/model"
	"github.com/amirhnajafiz/DJaaS/internal/model/relations"
)

type ClassRepository interface {
	Create(class *model.Class) error
	Delete(classID uint) error
	GetSingle(classID uint) (*model.Class, error)
	GetAll() ([]*model.Class, error)
	AddUser(userClass relations.UserClass) error
	RemoveUser(userClass relations.UserClass) error
	GetUsers(classID uint) ([]*model.User, error)
}
