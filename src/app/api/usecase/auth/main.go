package user

import (
	"errors"

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

func (u *Auth) Login(email string, password string) (*domain.Auth, error) {
	target := new(domain.Auth)

	u.repo.GetByFieldName(target, "email", email)

	if target.ID == 0 {
		return nil, errors.New("ERR_1001")
	}

	err := bcrypt.CompareHashAndPassword([]byte(target.EncryptedPassword), []byte(password))
	if err != nil {
		return nil, errors.New("ERR_1001")
	}

	return target, nil
}
