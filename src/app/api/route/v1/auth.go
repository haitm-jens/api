package v1

import (
	"pandog/app/api/handler"
	server "pandog/infra/local/server"
)

type (
	Auth struct {
		handler handler.Auth
	}
)

func NewAuth(h handler.Auth) Auth {
	return Auth{handler: h}
}

func (s Auth) Route() []server.Route {
	return []server.Route{
		server.Route{
			Method:   "POST",
			URL:      "/auth/login",
			Business: s.handler.Login,
		},
	}
}
