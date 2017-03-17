package auth

import (
	"service/models"
	"errors"
)

type UnactivatedUser struct {
	models.User
}

var unactivated_users map[string]models.User

func init() {
	unactivated_users = make(map[string]models.User)
}
func NewUser(pass string, usr models.User) error {
	if usr.FindByLogin() {
		return errors.New("This user is already registered")
	}
	if _, err := unactivated_users[pass]; err {
		return errors.New("Already in")
	}
	unactivated_users[pass] = usr
	return nil
}

func ActivateUser(pass string) error {
	if user, ok := unactivated_users[pass]; ok {
		delete(unactivated_users, pass)
		return user.Insert()
	} else {
		return errors.New("Wrong pass")
	}
}

var reset_passwords map[string]int

func RequestToResetPassword(pass string, usr models.User) error {
	if user, err := models.GetUserByLogin(usr.Login); err == nil {
		reset_passwords[pass] = user.Id
		return nil
	} else {
		return err
	}
}

func GetUserIdByPass(pass string) (int, bool) {
	id, err := reset_passwords[pass]
	return id, err
}