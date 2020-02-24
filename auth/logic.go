package auth

import (
	"jwtauth/auth/models"
	"time"
)

type authService struct {
	authRepo Repository
}

func (a authService) Login(username string) (*models.User, error) {
	return a.authRepo.Login(username)
}

func (a authService) Register(user *models.User) error {
	user.CreatedAt = time.Now().UTC().Unix()
	return a.authRepo.Register(user)
}

func NewAuthService(r Repository) Service{
	return &authService{
		authRepo: r,
	}
}
