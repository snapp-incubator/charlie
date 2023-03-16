package model

import (
	"github.com/amirhnajafiz/DJaaS/pkg/enum"
)

// User model.
// Username is unique.
// Password is for identification.
// Role is (Admin, Teacher, or Student)
type User struct {
	Base
	Username string    `json:"username"`
	Password string    `json:"password"`
	Role     enum.Role `json:"role"`
}
