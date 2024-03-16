package helper

import (
	"fmt"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(password, salt string) (string, error) {
	str := fmt.Sprintf("%s%s", password, salt)

	bytes, err := bcrypt.GenerateFromPassword([]byte(str), 14)
	return string(bytes), err
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandSalt(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func CheckPasswordHash(password, salt, hash string) bool {
	str := fmt.Sprintf("%s%s", password, salt)

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err == nil
}
