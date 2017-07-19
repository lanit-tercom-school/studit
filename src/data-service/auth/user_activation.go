package auth

import (
	"crypto/rand"
	"data-service/models"
	"fmt"
)

const (
	AvatarTemplatePath string = "http://tinygraphs.com/squares/"
	AvatarTemplateSize string = "100"
)

type UserActivationService struct {
}

type ActivationUser struct {
	Id       int
	Nickname string
}

func (s *UserActivationService) Activate(user *ActivationUser, reply *bool) error {
	avatar_seed := GenerateNewToken(6)
	color_str := GenerateRandomColor()
	u := models.User{
		Id:       user.Id,
		Nickname: user.Nickname,
		Avatar: fmt.Sprintf("%s%s?colors=%s&colors=%s&size=%s", AvatarTemplatePath, avatar_seed,
			color_str, "FFFFFF", AvatarTemplateSize),
	}
	return u.Insert()
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
