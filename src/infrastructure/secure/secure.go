package secure

import (
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Printf("error while hashing. Reason: %v\n", err)
	}
	return string(hash)
}
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Printf("error while comparing passwords, Reason: %v\n", err)
		return false
	}
	return true
}
