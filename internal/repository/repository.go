package repository

type Repository struct {
	Class    ClassRepository
	User     UserRepository
	Question QuestionRepository
	Submit   SubmitRepository
}

func New() *Repository {
	return &Repository{}
}
