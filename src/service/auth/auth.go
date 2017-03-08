package auth

import (
	"service/models"
	"errors"
	"crypto/rand"
)

type SessionStruct struct{
	Token 		string 		`json:"token"`
	UserId 		int 		`json:"-"`
	ExpiresIn	int64	`json:"expires_in"`
}

type Usr struct {
	Login string
	Password string
}

// TODO: добавить регистрацию
// TODO: добавить восстановление пароля
// TODO: добавить валидацию email (отправка проверочного кода на почту)
// basic http sign in with password and login. Func checks logic+password combo with combo in DB
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
	// TODO: UNcomment this on pub } else if user.Password != customStr(password).ToSHA1() {
	} else if user.Password != password { // TODO: comment this on pub
		return user, errors.New("Invalid login or password")
	} else {
		return user, nil  // all OK
	}
}

// Generates cryptographically secure token (random string)
func GenerateNewToken() string {
	// template string
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, 10) // TODO: change 10 to `n`
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}