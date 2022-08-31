package auth

import "github.com/firstaadi-dev/majoo-backend-test/domain"

type UserRepository interface {
	GetUser(username, password string) (*domain.User, error)
}
