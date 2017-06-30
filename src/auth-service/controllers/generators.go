package controllers

import "crypto/rand"

// Generates cryptographically secure token (random string)
func GenerateNewToken(count int) string {
	if count <= 0 {
		return ""
	}
	// template string
	const alphaNum= "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes= make([]byte, count)
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