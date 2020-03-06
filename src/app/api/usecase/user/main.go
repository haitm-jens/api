package user

import (
	"pandog/repo"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	repo repo.Repo
}

func NewUser(r repo.Repo) User {
	return User{repo: r}
}

func password(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}
