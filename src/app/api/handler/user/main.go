package user

import usecase "pandog/app/api/usecase/user"

type (
	User struct {
		Uc usecase.User
	}
)

func NewUser(u usecase.User) User {
	return User{Uc: u}
}
