package util

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func BcryptMake(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		logrus.Println(err)
	}
	return string(hash)
}

func BcryptMakeCheck(pwd []byte, hashedPwd string) bool {
	byteHash := []byte(hashedPwd)
	if err := bcrypt.CompareHashAndPassword(byteHash, pwd); err != nil {
		return false
	}
	return true
}
