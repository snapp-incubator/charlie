package model

// Class model.
// each class belongs to a user.
// class has a limit for the users.
type Class struct {
	Base
	UserID   uint `json:"user_id"`
	Capacity int  `json:"capacity"`
}
