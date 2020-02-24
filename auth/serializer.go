package auth

import "jwtauth/auth/models"

type Serializer interface {
	Encode(user *models.User) ([]byte, error)
	Decode(b []byte) (*models.User, error)
}
