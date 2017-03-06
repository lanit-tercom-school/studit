package auth

import (
	"service/models"
	"errors"
)

type Usr struct {
	Login string
	Password string
}

func TryToLogin(login, password string) (user models.User, err error) {
	// create default model
	user = &models.User{
		Login: login,
	}

	err = user.Read("login")
	if err != nil {
		return
	} else if user.Id < 1 {
		return user, errors.New("Bad user ID") // TODO: should be changed to "Invalid login or password"
	} else if user.Password != customStr(password).ToSHA1() {
		return user, errors.New("Invalid login or password")
	} else {
		return user, nil  // all OK
	}
}