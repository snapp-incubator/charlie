package repository

type Repository struct {
	User     UserRepository
	Question QuestionRepository
	Submit   SubmitRepository
}
