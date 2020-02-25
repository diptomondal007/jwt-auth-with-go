package auth

import "jwtauthwithgo/auth/models"

type Serializer interface {
	Encode(user *models.User) ([]byte, error)
	Decode(b []byte) (*models.User, error)
}
