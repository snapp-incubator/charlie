package relations

// UserClass relation.
type UserClass struct {
	Base
	UserID  uint `json:"user_id"`
	ClassID uint `json:"class_id"`
}
