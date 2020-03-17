package user

import (
	"pandog/domain"
	"pandog/repo"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	repo repo.Repo
}

func NewUser(r repo.Repo) User {
	return User{repo: r}
}

func (s *User) GetByID(id uint) *domain.User {
	target := new(domain.User)

	s.repo.GetByID(target, id)

	if target.ID == 0 {
		return nil
	}

	return target
}

func password(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}
