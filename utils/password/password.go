package pwd

import (
	"fmt"
	"math/rand"

	errs "github.com/kongsakchai/catopia-backend/domain/error"
	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(password string, salt string) (string, error) {
	str := fmt.Sprintf("%s:(-_-)7%s", password, salt)
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), 14)
	if err != nil {
		return "", errs.New(errs.ErrInternal, "Internal server error", err)
	}

	return string(bytes), err
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func Salt(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Compare(password, salt, hash string) bool {
	str := fmt.Sprintf("%s:(-_-)7%s", password, salt)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))

	return err == nil
}
