package auth

import "jwtauth/auth/models"

type Service interface {
	Login (username string) (*models.User, error)
	Register(user *models.User) error
}