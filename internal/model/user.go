package model

import (
	"github.com/amirhnajafiz/DJaaS/pkg/enum"
)

type User struct {
	Base
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Role     enum.Role `json:"role"`
}
