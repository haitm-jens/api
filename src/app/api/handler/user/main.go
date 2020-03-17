package user

import (
	usecase "pandog/app/api/usecase/user"
)

type (
	User struct {
		uc usecase.User
	}
)

func NewUser(u usecase.User) User {
	return User{uc: u}
}
