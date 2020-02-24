package auth

import "jwtauth/auth/models"

type Repository interface {
	Login(username string) (*models.User, error)
	Register(user *models.User) error
}
