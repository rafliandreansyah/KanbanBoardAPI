package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(p string) (string, error){
	pass := []byte(p)
	cost := 11

	password, err := bcrypt.GenerateFromPassword(pass, cost)
	if err != nil {
		return "", err
	}

	return string(password), nil

}