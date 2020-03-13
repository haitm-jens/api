package auth

import usecase "pandog/app/api/usecase/auth"

type (
	Auth struct {
		uc usecase.Auth
	}
)

func NewAuth(u usecase.Auth) Auth {
	return Auth{uc: u}
}
