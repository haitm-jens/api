package user

import (
	"fmt"
	"pandog/domain"
	"pandog/repo"

	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	repo repo.Repo
}

func NewAuth(repo repo.Repo) Auth {
	return Auth{repo: repo}
}

func (u *Auth) Login(email string, password string) *domain.Auth {
	target := new(domain.Auth)

	u.repo.GetByFieldName(target, "email", email)

	if target.ID == 0 {
		return nil
	}

	err := bcrypt.CompareHashAndPassword([]byte(target.EncryptedPassword), []byte(password))

	fmt.Println("err:", err)
	if err != nil {
		return nil
	}

	return target
}
