package json

import (
	"encoding/json"
	errs "github.com/pkg/errors"
	"jwtauth/auth/models"
)

type User struct{}

func (u *User) Encode(user *models.User) ([]byte, error) {
	raw, err := json.Marshal(user)
	if err != nil {
		return nil, errs.Wrap(err, "serializer.User.Encode")
	}
	return raw, nil
}

func (u *User) Decode(b []byte) (*models.User, error) {
	user := &models.User{}
	if err := json.Unmarshal(b, user); err != nil {
		return nil, errs.Wrap(err, "serializer.User.Decode")
	}
	return user, nil
}
