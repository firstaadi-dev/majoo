package auth

import "github.com/firstaadi-dev/majoo-backend-test/domain"

type UseCase interface {
	Login(username, password string) (string, error)
	ParseToken(accessToken string) (*domain.User, error)
}
