package repository

type Repository struct {
	User   UserRepository
	Submit SubmitRepository
}
