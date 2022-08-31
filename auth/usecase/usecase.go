package usecase

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/firstaadi-dev/majoo-backend-test/auth"
	"github.com/firstaadi-dev/majoo-backend-test/domain"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *domain.User `json:"user"`
}

type AuthUseCase struct {
	userRepo       auth.UserRepository
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthUsecase(
	userRepo auth.UserRepository,
	signingKey []byte,
	tokenTTL time.Duration) *AuthUseCase {
	return &AuthUseCase{
		userRepo:       userRepo,
		signingKey:     signingKey,
		expireDuration: time.Second * tokenTTL,
	}
}

func (a *AuthUseCase) Login(username, password string) (string, error) {

	user, err := a.userRepo.GetUser(username, password)
	if err != nil {
		return "", err
	}

	var claims = AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
		},
	}

	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.signingKey)
}
