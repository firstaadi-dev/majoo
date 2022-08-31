package auth

type UseCase interface {
	Login(username, password string) (string, error)
	// ParseToken(accessToken string) (*domain.User, error)
}
