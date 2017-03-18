package auth

import (
	"service/models"
	"errors"
)

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

// 					   ["login"]struct{"userId", "pass"}
var reset_passwords map[string]struct{int; string}

func RequestToResetPassword(pass string, usr models.User) error {
	if user, err := models.GetUserByLogin(usr.Login); err == nil {
		reset_passwords[user.Login] = struct {int; string}{user.Id, pass}
		return nil
	} else {
		return err
	}
}

func ResetPassword(login, pass, newPassword string) error {
	tuple, err := reset_passwords[login]
	if err && pass == tuple.string {
		u, err := models.GetUserById(tuple.int)
		if err == nil && u.Login == login {
			u.Password = newPassword
			u.Update()
			return nil
		} else {
			if err != nil {
				return err
			} else {
				return errors.New("Wrong login")
			}
		}
	} else {
		return errors.New("Wrong login or pass")
	}
}