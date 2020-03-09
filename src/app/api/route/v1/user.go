package v1

import (
	"pandog/app/api/handler"
	"pandog/infra/local/server"
)

type (
	User struct {
		handler handler.User
	}
)

func NewUser(h handler.User) User {
	return User{handler: h}
}

func (s User) Route() []server.Route {
	return []server.Route{
		server.Route{
			Method:   "GET",
			URL:      "users/me",
			Business: s.handler.Get,
		},
	}
}
