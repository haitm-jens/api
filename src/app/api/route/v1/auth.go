package v1

import (
	"pandog/app/api/handler"
	usecase "pandog/app/api/usecase/user"
	server "pandog/infra/server"
	database "pandog/interface/infra/db"
)

type (
	auth struct {
		handler *handler.Auth
	}
)

func NewAuth(db database.MySQL) *auth {
	uc := usecase.NewUser(db)
	hl := handler.NewAuth(*uc)
	return &auth{handler: hl}
}

func (s *auth) Route() []server.Route {
	return []server.Route{
		server.Route{
			Method:   "POST",
			URL:      "auth",
			Business: s.handler.Login,
		},
	}
}
