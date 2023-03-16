package repository

type Repository struct {
	Class    ClassRepository
	User     UserRepository
	Question QuestionRepository
	Submit   SubmitRepository
}
