package auth

import (
	"service/models"
	"errors"
	"crypto/rand"
)

// TODO: rename
type UserAndToken struct {
	Token           string              `json:"bearer_token"`
	User            models.MainUserInfo `json:"user"`
	ExpiresIn       string              `json:"exp"`
	PermissionLevel int                 `json:"perm_lvl"`
}

type Usr struct {
    Login string    `json:"login"`
    Password string `json:"password"`
}

const MaxPermissionLevel int = 2

// basic http sign in with password and login. Func checks login+password combination with same combination in DB
func TryToLogin(login, password string) (user models.User, err error) {
	// create default model
	user = models.User{
		Login: login,
	}

	err = user.Read("login")
	if err != nil {
		return user, errors.New("Can't find User with this login (dev)") // TODO: should be changed to "Invalid login or password"
	} else if user.Id < 1 {
		return user, errors.New("Bad user ID (dev)") // TODO: should be changed to "Invalid login or password"
	} else if user.Password != CustomStr(password).ToSHA1() { // TODO: UNComment this on pub
	//} else if user.Password != password { // TODO: comment this on pub
		return user, errors.New("Invalid login or password")
	} else {
		return user, nil  // all OK
	}
}

func TryToChangePassword(user_id int, password string) (user *models.User, err error) {
	// create default model
	user, err = models.GetUserById(user_id)

	if err != nil {
		return user, errors.New("Can't find User with this login (dev)") // TODO: should be changed to "Invalid login or password"
	} else if user.Id < 1 {
		return user, errors.New("Bad user ID (dev)") // TODO: should be changed to "Invalid login or password"
	} else if user.Password != CustomStr(password).ToSHA1() { // TODO: UNComment this on pub
	//} else if user.Password != password { // TODO: comment this on pub
		return user, errors.New("Invalid login or password")
	} else {
		return user, nil  // all OK
	}
}

func ChangePasswordForUser(user *models.User) error {
	user.Password = CustomStr(user.Password).ToSHA1()
	return user.Update()
}

// Generates cryptographically secure token (random string)
func GenerateNewToken(count int) string {
	if count <= 0 {
		return ""
	}
	// template string
	const alphaNum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, count)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphaNum[b%byte(len(alphaNum))]
	}
	return string(bytes)
}

func GenerateRandomColor() string {
	// template string
	const alphaNum = "0123456789ABCDEF"
	var bytes = make([]byte, 6)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphaNum[b%byte(len(alphaNum))]
	}
	return string(bytes)
}

